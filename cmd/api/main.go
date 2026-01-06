package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn         string
		maxConns    int
		minConns    int
		maxIdleTime time.Duration
	}
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config
	if err := godotenv.Load(); err != nil {
		fmt.Fprintln(os.Stderr, "Error loading .env file")
	}

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("PRISM_STACK_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxConns, "db-max-conns", 50, "PostgerSQL max connections")
	flag.IntVar(&cfg.db.minConns, "db-min-conns", 2, "PostgreSQL min connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err = srv.ListenAndServe()
	slog.Error(err.Error())
	os.Exit(1)
}

func openDB(cfg config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(cfg.db.maxConns)
	config.MinConns = int32(cfg.db.minConns)
	config.MaxConnIdleTime = cfg.db.maxIdleTime

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.Ping(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

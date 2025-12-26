package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/stocks", app.createStockHandler)
	router.HandlerFunc(http.MethodGet, "/v1/stocks/:id", app.showStockHandler)

	return router
}

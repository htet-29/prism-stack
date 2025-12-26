package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/htet-29/prism-stack/internal/data"
)

func (app *application) createStockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new stock")
}

func (app *application) showStockHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	stock := data.Stock{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ItemName:  "Beer",
		SKU:       "BA123",
		Category:  nil,
		UnitPrice: 0,
		Quantity:  0,
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"stock": stock}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

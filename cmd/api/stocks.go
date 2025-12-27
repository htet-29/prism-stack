package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/htet-29/prism-stack/internal/data"
)

func (app *application) createStockHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ItemName  string   `json:"item_name"`
		SKU       string   `json:"sku"`
		Category  []string `json:"category"`
		UnitPrice int64    `json:"unit_price"`
		Quantity  int64    `json:"quantity"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showStockHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	stock := data.Stock{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ItemName:  "Beer",
		SKU:       "BA123",
		Category:  []string{"Liquor"},
		UnitPrice: 100,
		Quantity:  10,
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"stock": stock}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

package main

import (
	"fmt"
	"net/http"
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

	fmt.Fprintf(w, "showing the details of stock: %d\n", id)
}

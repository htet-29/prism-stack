package main

import (
	"fmt"
	"net/http"
)

// recoverPanic is the middleware that recovers from panics and send error response to client
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// a deferred function (which will always be run in the event of a panic)
		defer func() {
			if pv := recover(); pv != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%v", pv))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

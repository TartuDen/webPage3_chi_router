package main

import (
	"fmt"
	"net/http"
)

// WriteToConsole is a middleware that logs a message to the console
// each time an HTTP request is received. It takes an `http.Handler` as
// an argument, allowing it to be chained with other middleware or handlers.
func WriteToConsole(next http.Handler) http.Handler {
	// Return an http.Handler that logs a message to the console and then
	// calls the ServeHTTP method of the next handler in the chain.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log a message to the console indicating that the page has been hit.
		fmt.Println("Hit the page!")

		// Call the ServeHTTP method of the next handler in the chain
		// to continue processing the HTTP request.
		next.ServeHTTP(w, r)
	})
}


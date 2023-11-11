package main

import (
	"net/http"

	"github.com/TartuDen/webPage3_chi_router/pkg/config"
	"github.com/TartuDen/webPage3_chi_router/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(a *config.AppConfig) http.Handler {
	// Create a new instance of the chi router.
	mux := chi.NewRouter()

	// Use the Recoverer middleware from chi.
	// This middleware recovers and logs panics, preventing the application from crashing.
	mux.Use(middleware.Recoverer)

	// Use the NoSurf middleware with the chi router.
	// This adds CSRF protection to all routes registered with this router.
	mux.Use(NoSurf)

	// Add a route for HTTP GET requests to the root path ("/").
	// Associate the MainHandler function from the handler.Repo struct with this route.
	mux.Get("/", handler.Repo.MainHandler)

	// Add a route for HTTP GET requests to the "/about" path.
	// Associate the AboutHandler function from the handler.Repo struct with this route.
	mux.Get("/about", handler.Repo.AboutHandler)

	// Return the configured chi router.
	return mux
}

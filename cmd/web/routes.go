package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/housecham/FirstWebApp/pkg/config"
	"github.com/housecham/FirstWebApp/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	// Routes setup
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))       // Home()
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About)) // About()

	return mux
}

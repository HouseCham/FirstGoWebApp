package main

import (
	"net/http"

	//"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
	"github.com/housecham/FirstWebApp/pkg/config"
	"github.com/housecham/FirstWebApp/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// CHI middlewares
	mux.Use(middleware.Recoverer)
	
	// using my own middlewares
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

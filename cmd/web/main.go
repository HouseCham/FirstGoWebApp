package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/housecham/FirstWebApp/pkg/config"
	"github.com/housecham/FirstWebApp/pkg/handlers"
	"github.com/housecham/FirstWebApp/pkg/render"
)

const portNumber string = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Share config file with render.go
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port http://localhost%s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
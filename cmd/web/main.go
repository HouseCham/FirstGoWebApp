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

	fmt.Printf("Starting application on port http://localhost%s", portNumber)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}

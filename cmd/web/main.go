package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/housecham/FirstWebApp/pkg/config"
	"github.com/housecham/FirstWebApp/pkg/handlers"
	"github.com/housecham/FirstWebApp/pkg/render"
)

const portNumber string = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// set to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour //This session will last for 24 hrs
	session.Cookie.Persist = true     // This cookie will persist after the window is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // To insist the cookie to be encrypted -> to use only https... in production set to true

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	// Share config file with handlers.go
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

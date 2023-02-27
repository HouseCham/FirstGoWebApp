package main

import (
	"fmt"
	"net/http"
	"github.com/housecham/FirstWebApp/pkg/handlers"
)

const portNumber string = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port http://localhost%s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
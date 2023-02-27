package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// Render templates using HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.html")
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
package handlers

import (
	"net/http"

	"github.com/housecham/FirstWebApp/pkg/config"
	"github.com/housecham/FirstWebApp/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.html")
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.html")
}
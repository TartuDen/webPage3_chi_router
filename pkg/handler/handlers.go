package handler

import (
	"net/http"

	"github.com/TartuDen/webPage3_chi_router/pkg/config"
	"github.com/TartuDen/webPage3_chi_router/pkg/models"
	"github.com/TartuDen/webPage3_chi_router/pkg/renderer"
)

// TemplateData holds data sent from handlers to templates

// Repo the repository used by the handlers
var Repo *Repository

// Repositroy is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// MainHandler is a method of the Repository struct that handles requests to the main page.
// It renders the "home.page.html" template to the provided HTTP response writer.
func (m *Repository) MainHandler(w http.ResponseWriter, r *http.Request) {
	remoteIP:=r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)
	renderer.RendererTemplate(w, "home.page.html", &models.TemplateData{})
}

// AboutHandler is a method of the Repository struct that handles requests to the about page.
// It renders the "about.page.html" template to the provided HTTP response writer.
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap:=make(map[string]string)
	stringMap["test"]="this is test data!"
	remoteIP:= m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"]=remoteIP

	//send data to the template
	renderer.RendererTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

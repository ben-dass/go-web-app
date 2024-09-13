package handlers

import (
    "net/http"
    
    "go-web-app-p1/pkg/config"
    "go-web-app-p1/pkg/models"
    "go-web-app-p1/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
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

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    remoteIP := r.RemoteAddr
    m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
    
    render.TemplateRender(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    stringMap := make(map[string]string)
    stringMap["test"] = "Hello again."
    
    remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
    stringMap["remote_ip"] = remoteIP
    
    render.TemplateRender(w, "about.page.tmpl", &models.TemplateData{
        StringMap: stringMap,
    })
}

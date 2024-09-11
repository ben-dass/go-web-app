package handlers

import (
    "net/http"
    
    "go-web-app-p1/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
    render.Template(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
    render.Template(w, "about.page.tmpl")
}

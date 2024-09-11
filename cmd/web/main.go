package main

import (
    "fmt"
    "log"
    "net/http"
    
    "go-web-app-p1/pkg/config"
    "go-web-app-p1/pkg/handlers"
    "go-web-app-p1/pkg/render"
)

const webPort = ":8080"

func main() {
    var app config.AppConfig
    
    tc, err := render.CreateTemplateCache()
    if err != nil {
        log.Fatal("cannot create template cache")
    }
    
    app.TemplateCache = tc
    
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)
    
    fmt.Printf("Starting server at port %s\n", webPort)
    _ = http.ListenAndServe(webPort, nil)
}

package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    
    "github.com/alexedwards/scs/v2"
    "github.com/ben-dass/go-web-app/pkg/config"
    "github.com/ben-dass/go-web-app/pkg/handlers"
    "github.com/ben-dass/go-web-app/pkg/render"
)

const webPort = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
    tc, err := render.CreateTemplateCache()
    if err != nil {
        log.Fatal("cannot create template cache")
    }
    
    app.TemplateCache = tc
    app.UseCache = false
    app.InProduction = false
    
    session = scs.New()
    session.Lifetime = 24 * time.Hour
    session.Cookie.Persist = true
    session.Cookie.SameSite = http.SameSiteLaxMode
    session.Cookie.Secure = app.InProduction
    
    app.Session = session
    
    repo := handlers.NewRepo(&app)
    handlers.NewHandlers(repo)
    
    render.NewTemplates(&app)
    
    srv := &http.Server{
        Addr:    webPort,
        Handler: routes(&app),
    }
    
    fmt.Printf("Starting server at port %s\n", webPort)
    err = srv.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}

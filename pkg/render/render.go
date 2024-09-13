package render

import (
    "bytes"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
    
    "go-web-app-p1/pkg/config"
    "go-web-app-p1/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
    app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
    return td
}

func TemplateRender(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
    var tc map[string]*template.Template
    
    if app.UseCache {
        tc = app.TemplateCache
    } else {
        tc, _ = CreateTemplateCache()
    }
    
    t, ok := tc[tmpl]
    if !ok {
        log.Fatal("Error loading template:", tmpl)
    }
    
    buff := new(bytes.Buffer)
    td = AddDefaultData(td)
    
    err := t.Execute(buff, td)
    if err != nil {
        log.Println("Error executing template:", err)
    }
    
    _, err = buff.WriteTo(w)
    if err != nil {
        log.Println("Error writing:", err)
    }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
    tc := map[string]*template.Template{}
    
    pages, err := filepath.Glob("./templates/*.page.tmpl")
    if err != nil {
        return nil, err
    }
    
    for _, page := range pages {
        name := filepath.Base(page)
        
        templateSet, err := template.New(name).ParseFiles(page)
        if err != nil {
            return tc, err
        }
        
        matches, err := filepath.Glob("./templates/*.layout.tmpl")
        if err != nil {
            return tc, err
        }
        
        if len(matches) > 0 {
            templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
            if err != nil {
                return tc, err
            }
        }
        
        tc[name] = templateSet
    }
    
    return tc, nil
}

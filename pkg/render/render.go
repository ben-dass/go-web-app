package render

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
    var tmpl *template.Template
    var err error
    
    _, inMap := templateCache[t]
    if !inMap {
        log.Println("creating template and adding to cache")
        err = createTemplateCache(t)
        if err != nil {
            fmt.Println(err)
        }
    } else {
        log.Println("using cached template")
        tmpl = templateCache[t]
    }
    
    tmpl = templateCache[t]
    
    err = tmpl.Execute(w, nil)
    if err != nil {
        fmt.Println(err)
    }
}

func createTemplateCache(t string) error {
    templates := []string{
        fmt.Sprintf("./templates/%s", t),
        "./templates/base.layout.tmpl",
    }
    
    tmpl, err := template.ParseFiles(templates...)
    if err != nil {
        return err
    }
    
    templateCache[t] = tmpl
    
    return nil
}

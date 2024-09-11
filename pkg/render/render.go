package render

import (
    "bytes"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

func Template(w http.ResponseWriter, tmpl string) {
    tmplCache, err := CreateTemplateCache()
    if err != nil {
        log.Fatal("Error creating template cache:", err)
    }
    
    t, ok := tmplCache[tmpl]
    if !ok {
        log.Fatal("Error loading template:", tmpl)
    }
    
    buff := new(bytes.Buffer)
    
    err = t.Execute(buff, nil)
    if err != nil {
        log.Println("Error executing template:", err)
    }
    
    _, err = buff.WriteTo(w)
    if err != nil {
        log.Println("Error writing:", err)
    }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
    templateCache := map[string]*template.Template{}
    
    pages, err := filepath.Glob("./templates/*.page.tmpl")
    if err != nil {
        return nil, err
    }
    
    for _, page := range pages {
        name := filepath.Base(page)
        
        templateSet, err := template.New(name).ParseFiles(page)
        if err != nil {
            return templateCache, err
        }
        
        matches, err := filepath.Glob("./templates/*.layout.tmpl")
        if err != nil {
            return templateCache, err
        }
        
        if len(matches) > 0 {
            templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
            if err != nil {
                return templateCache, err
            }
        }
        
        templateCache[name] = templateSet
    }
    
    fmt.Println(templateCache)
    return templateCache, nil
}

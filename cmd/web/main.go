package main

import (
    "fmt"
    "net/http"
    
    "go-web-app-p1/pkg/handlers"
)

const webPort = ":8080"

func main() {
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)
    
    fmt.Printf("Starting server at port %s\n", webPort)
    _ = http.ListenAndServe(webPort, nil)
}

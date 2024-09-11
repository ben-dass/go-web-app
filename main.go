package main

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprintf(w, "This is the Home page.")
    if err != nil {
        fmt.Println(err)
    }
}

func About(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprintf(w, "This is the About page.")
    if err != nil {
        fmt.Println(err)
    }
}

func AddValues(x, y int) int {
    return x + y
}

func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about", About)
    
    _ = http.ListenAndServe(":8080", nil)
}

package main

import (
        "log"
        "net/http"
       )

func main() {

    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
                http.ServeFile(w, r, r.URL.Path[1:])
    })

    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/main.html")
}

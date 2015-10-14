package main

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/data/[[:alpha:]]", dataHandler)
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("home handler serving request")
	log.Println(r)
	http.ServeFile(w, r, "static/main.html")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("data handler serving request")
	log.Println(r)
	re := regexp.MustCompile(".*/[[:alpha:]]$")
	path := r.URL.Path
	food := re.FindString(path)
	io.WriteString(w, "Hello "+food)
}

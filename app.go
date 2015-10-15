package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Data struct {
	ID               int
	Name             string
	GallonPerCalorie float64
}

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/data/{food}", dataHandler)
	mux.HandleFunc("/static/{folder}/{file}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	port := ":3000"
	log.Println("Listening on " + port)
	http.ListenAndServe(port, mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/main.html")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := Data{
		ID:               1,
		Name:             vars["food"],
		GallonPerCalorie: 1.34,
	}
	b, _ := json.Marshal(data)
	io.WriteString(w, string(b))
}

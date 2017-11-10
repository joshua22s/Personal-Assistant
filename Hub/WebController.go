package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startWebServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))
	log.Fatal(http.ListenAndServe(":8000", r))
}

//url: /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("web/html/index.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, nil)
	}
}

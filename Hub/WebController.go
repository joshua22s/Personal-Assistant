package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var tPath = "./temps/"
var dPath = "./data/"

var templates = template.Must(template.ParseFiles(tPath+"home.html", dPath+"header.html", dPath+"head.html", dPath+"footer.html"))

//var templateDirs = []string{"temps", "data"}

func rootHandler(wr http.ResponseWriter, req *http.Request) {
	title := "hello"

	headFile, headErr := ioutil.ReadFile(dPath + "head.html")
	headerFile, headerErr := ioutil.ReadFile(dPath + "header.html")
	footerFile, footErr := ioutil.ReadFile(dPath + "footer.html")

	if headErr != nil || headerErr != nil || footErr != nil {
		http.Error(wr, headErr.Error(), http.StatusInternalServerError)
		http.Error(wr, headerErr.Error(), http.StatusInternalServerError)
		http.Error(wr, footErr.Error(), http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"title":  title,
		"header": string(headerFile),
		"head":   string(headFile),
		"footer": string(footerFile),
	}

	err := templates.ExecuteTemplate(wr, "homeHTML", data)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

func startWebServer() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

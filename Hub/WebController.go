package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var tPath = "./temps/"
var dPath = "./data/"

var templates = template.Must(template.ParseFiles(tPath+"home.html", tPath+"login.html", dPath+"header.html", dPath+"loginheader.html", dPath+"head.html", dPath+"footer.html"))

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

func loginHandler(wr http.ResponseWriter, req *http.Request) {
	title := "Login"

	headFile, headErr := ioutil.ReadFile(dPath + "head.html")
	headerFile, headerErr := ioutil.ReadFile(dPath + "loginheader.html")
	footerFile, footErr := ioutil.ReadFile(dPath + "footer.html")

	if headErr != nil || headerErr != nil || footErr != nil {
		http.Error(wr, headErr.Error(), http.StatusInternalServerError)
		http.Error(wr, headerErr.Error(), http.StatusInternalServerError)
		http.Error(wr, footErr.Error(), http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"title":       title,
		"head":        string(headFile),
		"loginheader": string(headerFile),
		"footer":      string(footerFile),
	}

	err := templates.ExecuteTemplate(wr, "loginHTML", data)

	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

func startWebServer() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//	fs := http.FileServer(http.Dir("data"))
	//	http.Handle("/css/", fs)
	//	http.Handle("/js/", fs)
	//	http.Handle("/fonts/", fs)
	//	http.Handle("/img/", fs)
	http.ListenAndServe(":8080", nil)
}

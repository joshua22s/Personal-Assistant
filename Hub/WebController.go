package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/index.html")
		t.Execute(w, nil)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("web/todos.html")
		if err != nil {
			fmt.Println(err)
		}
		model := MorningTodoModel{getAllUserMorningTodos(1)}
		fmt.Print(model)
		t.Execute(w, model)
	}
}

func startWebServer() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/todos", todoHandler)
	//	http.HandleFunc("/login", loginHandler)
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)
	http.Handle("/fonts/", fs)
	http.Handle("/img/", fs)
	fmt.Println("Started webserver, listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

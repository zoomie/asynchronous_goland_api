package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, r.Form)
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func main() {
	fmt.Println("up and running")
	server := http.Server{
		Addr : "127.0.0.1:8000",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/", form) 
	server.ListenAndServe()
}
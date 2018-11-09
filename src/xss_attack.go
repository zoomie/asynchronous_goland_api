package main

import (
	"net/http"
	"html/template"
)

func process(w http.ResponseWriter,r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr : "8000"
	}
	http.HandleFunc("/form", form) 
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
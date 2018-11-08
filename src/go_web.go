package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.parse("template_name.html")
	// t.Execute(w, "template_name.html")
	fmt.Fprintf(w, "hello!")
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
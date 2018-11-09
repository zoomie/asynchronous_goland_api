package main

import (
	"fmt"
	"net/http"
	"html/template"
	"math/rand"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	rand.Seed(time.Now().Unix())
	cond := rand.Intn(10) > 5
	t.Execute(w, cond)
}

func main() {
	fmt.Println("Starting up the server...")
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
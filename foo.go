package main

import (
    "encoding/json"
    "net/http"
    "sync"  
)

type MyStruct struct {
    A, B, Total int64
}

func (s *MyStruct) add() {
    s.Total = s.A + s.B
}

func process(w http.ResponseWriter, cr chan *http.Request) {
    r := <- cr
    var s MyStruct
    json.NewDecoder(r.Body).Decode(&s)
    s.add()
    json.NewEncoder(w).Encode(s)
}

func handler(w http.ResponseWriter, r *http.Request) {  
    cr := make(chan *http.Request, 1)
    cr <- r
    var pleasewait sync.WaitGroup
    pleasewait.Add(1)

    go func() {
        defer pleasewait.Done()
        process(w, cr) // doesn't work; no response :-(
    }()
    // process(w, cr) // works, but blank response :-(
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
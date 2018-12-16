package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    // METHOD 1
    logs := make(chan string)
    go logLogs(logs)
    handleHello := makeHello(logs)

    // METHOD 2
    passer := &DataPasser{logs: make(chan string)}
    go passer.log()

    http.HandleFunc("/1", handleHello)
    http.HandleFunc("/2", passer.handleHello)
    http.ListenAndServe(":5000", nil)
}

// METHOD 1

func makeHello(logger chan string) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        logger <- r.Host
        io.WriteString(w, "Hello world!")
    }
}

func logLogs(logger chan string) {
    for item := range logger {
        fmt.Println("1. Item", item)
    }
}

// METHOD 2

type DataPasser struct {
    logs chan string
}

func (p *DataPasser) handleHello(w http.ResponseWriter, r *http.Request) {
    p.logs <- r.URL.String()
    io.WriteString(w, "Hello world")
}

func (p *DataPasser) log() {
    for item := range p.logs {
        fmt.Println("2. Item", item)
    }
}
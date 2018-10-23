package main

import (
	"fmt"
	"github.com/gorilla/mux"
)

func SubtractOneFromLength(slice []byte) []byte {
    slice = slice[0 : len(slice)-1]
    return slice
}

func main() {
	var buffer [256]byte
	slice := buffer[50:100]
    fmt.Println("Before: len(slice) =", len(slice))
    slice = SubtractOneFromLength(slice)
    fmt.Println("After:  len(slice) =", len(slice))
    // fmt.Println("After:  len(newSlice) =", len(newSlice))
}

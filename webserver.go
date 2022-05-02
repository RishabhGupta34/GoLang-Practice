package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
	io.WriteString(w, "Hello world!")
}
func sec(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello!")
}

func main() {
	fs := http.FileServer(http.Dir("images"))
	// http.HandleFunc("/", hello)
	http.Handle("/", fs)
	http.HandleFunc("/sec/h", sec)
	http.ListenAndServe(":8000", nil)
}

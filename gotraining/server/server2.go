package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type book struct {
	Name   string
	Author string
}

func Capital(w http.ResponseWriter, r *http.Request) {
	bt, _ := ioutil.ReadFile("serialdatabook.json")
	io.WriteString(w, strings.ToUpper(string(bt)))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/capital", Capital).Methods("GET")
	http.ListenAndServe(":8080", r)
}

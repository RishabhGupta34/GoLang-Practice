package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type book struct {
	Name   string
	Author string
}

func reader(w http.ResponseWriter, r *http.Request) {
	bt, _ := ioutil.ReadFile("serialdatabook.json")
	var data []book
	json.Unmarshal(bt, &data)
	w.Write(bt)
}
func capitalizer(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/capital")
	fmt.Println(err)
	inp, err := ioutil.ReadAll(resp.Body)
	fmt.Println(err)
	w.Write(inp)
}

func writer(w http.ResponseWriter, r *http.Request) {
	bt, _ := ioutil.ReadFile("serialdatabook.json")
	var data []book
	json.Unmarshal(bt, &data)
	inp, _ := ioutil.ReadAll(r.Body)
	var newdata book
	json.Unmarshal(inp, &newdata)
	data = append(data, newdata)
	bt, _ = json.Marshal(data)
	err := ioutil.WriteFile("serialdatabook.json", bt, 0766)
	fmt.Println(err)
	// w.Write([]byte("Done"))
	io.WriteString(w, "abcd")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/book", reader).Methods("GET")
	r.HandleFunc("/book", writer).Methods("POST")
	r.HandleFunc("/capital", capitalizer).Methods("GET")
	http.ListenAndServe(":8000", r)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type book struct {
	Name   string
	Author string
}

type Librarybook struct {
	Title      string `json:"Name"`
	AuthorName string `json:"Author"`
}

var books = []book{
	{"Harry Potter", "J.K. Rowling"},
	{"Harry Potter", "J.K. Rowling"},
}

var defstring = `{{range .}}
the book {{.Name}}
was written by {{.Author}}
{{end}}
`

func main() {
	t := template.New("Book template")
	t, _ = t.Parse(defstring)
	// b := book{"Harry Potter", "J.K. Rowling"}
	// t.Execute(os.Stdout, b)
	err := t.Execute(os.Stdout, books)
	fmt.Println(err, "\n")
	bt, err := json.Marshal(books)
	fmt.Println(err, "\n")
	err = ioutil.WriteFile("serialdatabook.json", bt, 0766)
	bt, _ = ioutil.ReadFile("serialdatabook.json")
	fmt.Println(string(bt))
	var lb []Librarybook
	json.Unmarshal(bt, &lb)
	fmt.Printf("%+v\n", lb)
}

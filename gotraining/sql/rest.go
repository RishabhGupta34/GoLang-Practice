package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type book struct {
	Id     int
	Name   string
	Author string
}

var db *sql.DB
var err error

func reader(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM user2_books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	bks := make([]book, 0)
	for rows.Next() {
		var bk book
		err = rows.Scan(&bk.Id, &bk.Name, &bk.Author)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	bt, err := json.Marshal(bks)
	if err != nil {
		log.Fatal(err)
	}
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
	inp, _ := ioutil.ReadAll(r.Body)
	var newdata book
	json.Unmarshal(inp, &newdata)
	s := fmt.Sprintf(`insert into user2_books(id,title,author) values('%v', '%s', '%s')`, newdata.Id, newdata.Name, newdata.Author)
	_, err = db.Exec(s)
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(w, "abcd")
}
func main() {

	db, err = sql.Open("mysql", "user2:user2@tcp(34.93.97.32:3306)/booksdb")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	create table if not exists user2_books (
		id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
		title varchar(100),
		author varchar(100)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/book", reader).Methods("GET")
	r.HandleFunc("/book", writer).Methods("POST")
	r.HandleFunc("/capital", capitalizer).Methods("GET")
	http.ListenAndServe(":8000", r)
}

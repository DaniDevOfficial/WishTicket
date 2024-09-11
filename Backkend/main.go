package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/wishticket")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequest()
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World CurrentTime: %s", time.Now())
}

func handleRequest() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server is listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

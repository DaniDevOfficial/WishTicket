package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"wishticket/modules/dev"
	"wishticket/modules/ticket"
	"wishticket/modules/user"

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

	initDB()

	router := http.NewServeMux()

	user.RegisterUserRoute(router, db)
	ticket.RegisterTicketRoute(router, db)
	dev.RegisterTicketRoute(router, db)
	fmt.Println("Server is listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

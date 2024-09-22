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

	handler := corsMiddleware(router)

	fmt.Println("Server is listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe("localhost:8000", handler))
}

// corsMiddleware sets the CORS headers to allow all origins
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// goofy ah request for checking if the server supports CORS or sum bs
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

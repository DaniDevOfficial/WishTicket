package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

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

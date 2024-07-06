package main

import (
	"fmt"
	"net/http"
)

func testHomePage (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the beginning of the Blog")
}

func handleRequests() {
	http.HandleFunc ("/", testHomePage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}


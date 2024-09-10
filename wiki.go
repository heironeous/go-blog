package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", pageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

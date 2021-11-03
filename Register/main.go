package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	server := http.ListenAndServe(":4444", router)
	log.Fatal(server)
}

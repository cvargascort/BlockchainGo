package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	server := http.ListenAndServe(":8003", router)
	log.Fatal(server)
}

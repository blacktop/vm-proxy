package main

import (
	"log"
	"net/http"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

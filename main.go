package main

import (
	"log"
	"net/http"

	"example.com/library"
)

func main() {
	http.HandleFunc("/", library.HelloWorld)
	log.Fatal(http.ListenAndServe(":8545", nil))
}

package library

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Balaji Pachai, is the wisest blockchain")
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a GET request\n")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a POST request\n")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a DELETE request\n")
}
func Run(address string) {
	r := mux.NewRouter()

	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")
	r.HandleFunc("/hello", helloWorld).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server started on port:", address)
	log.Fatal(http.ListenAndServe(":8545", nil))
}

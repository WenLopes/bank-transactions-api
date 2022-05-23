package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld).Methods("GET")
	fmt.Printf("Escutando na porta LOCAL 16091 🏆 \n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	helloWorld := new(struct {
		Message string `json:"message"`
	})
	helloWorld.Message = "Hello world!"

	json.NewEncoder(w).Encode(helloWorld)
}

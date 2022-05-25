package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Printf("Escutando na porta LOCAL 16091 🏆 \n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

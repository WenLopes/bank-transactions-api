package handlers

import (
	"fmt"
	"net/http"

	"github.com/WenLopes/bank-transactions-api/domain"
	"github.com/gorilla/mux"
)

func NewResetHandler(router *mux.Router, accountRepo domain.AccountRepository) {
	router.HandleFunc("/reset", reset(accountRepo)).Methods("GET").Name("reset")
}

func reset(accountRepo domain.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("todo: resetar estado")
	}
}

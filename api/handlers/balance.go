package handlers

import (
	"fmt"
	"net/http"

	"github.com/WenLopes/bank-transactions-api/app/account"
	"github.com/gorilla/mux"
)

func NewBalanceHandler(router *mux.Router, accountService account.UseCase) {
	router.HandleFunc("/balance", getBalance(accountService)).Methods("GET").Name("getBalance")
}

func getBalance(accountService account.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("todo: obter saldo")
	}
}

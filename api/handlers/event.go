package handlers

import (
	"fmt"
	"net/http"

	"github.com/WenLopes/bank-transactions-api/app/account"
	"github.com/gorilla/mux"
)

func NewEventHandler(router *mux.Router, accountService account.UseCase) {
	router.HandleFunc("/event", newEvent(accountService)).Methods("POST").Name("newEvent")
}

func newEvent(accountService account.UseCase) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("todo: executar evento")
	}
}

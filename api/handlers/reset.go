package handlers

import (
	"net/http"

	"github.com/WenLopes/bank-transactions-api/api/responses"
	"github.com/WenLopes/bank-transactions-api/app/account"
	"github.com/gorilla/mux"
)

func NewResetHandler(router *mux.Router, accountService account.UseCase) {
	router.HandleFunc("/reset", reset(accountService)).Methods("GET").Name("reset")
}

func reset(accountService account.UseCase) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		accountService.DeleteAll()
		responses.JSON(writer, http.StatusOK, "OK")
	}
}

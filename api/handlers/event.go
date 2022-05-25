package handlers

import (
	"fmt"
	"net/http"

	"github.com/WenLopes/bank-transactions-api/api/requests"
	"github.com/WenLopes/bank-transactions-api/api/responses"
	"github.com/WenLopes/bank-transactions-api/app/account"
	"github.com/gorilla/mux"
)

var eventHandleable = map[string]func(w http.ResponseWriter, event requests.EventRequest, accountService account.UseCase){
	"deposit":  handleDeposit,
	"transfer": handleTransfer,
	"withdraw": handleWithDraw,
}

func NewEventHandler(router *mux.Router, accountService account.UseCase) {
	router.HandleFunc("/event", newEvent(accountService)).Methods("POST").Name("newEvent")
}

func newEvent(accountService account.UseCase) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		event, err := requests.MapRequestToEvent(request)
		if err != nil {
			responses.Error(writer, http.StatusUnprocessableEntity, err) // NÃO DEVOLVER MENSAGEM DE ERRO PARA O USUÁRIO
			return
		}

		eventHandleable[event.EventType](writer, event, accountService)
	}
}

func handleDeposit(
	writer http.ResponseWriter,
	event requests.EventRequest,
	accountService account.UseCase,
) {
	fmt.Println("handle deposit")
}

func handleWithDraw(
	writer http.ResponseWriter,
	event requests.EventRequest,
	accountService account.UseCase,
) {
	fmt.Println("handle withdraw")
}

func handleTransfer(
	writer http.ResponseWriter,
	event requests.EventRequest,
	accountService account.UseCase,
) {
	fmt.Println("handle transfer")
}

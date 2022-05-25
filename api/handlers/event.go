package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/WenLopes/bank-transactions-api/api/presenters"
	"github.com/WenLopes/bank-transactions-api/api/requests"
	"github.com/WenLopes/bank-transactions-api/api/responses"
	"github.com/WenLopes/bank-transactions-api/api/validators"
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

	_, err := validators.ValidateDeposit(event)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	accountId, err := strconv.Atoi(event.Destination)

	if err != nil {
		fmt.Println(err) // Logar erro
		responses.Error(writer, http.StatusInternalServerError, errors.New("não foi possível concluir a operação"))
		return
	}

	account, err := accountService.ExecuteDeposit(accountId, event.Amount)

	if err != nil {
		fmt.Println(err) //logar erro
		responses.Error(writer, http.StatusInternalServerError, errors.New("não foi possível concluir a operação"))
		return
	}

	presenter := presenters.NewDepositPresenter(account.Id, account.Balance)
	responses.JSON(writer, http.StatusOK, presenter)
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

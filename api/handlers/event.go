package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/WenLopes/bank-transactions-api/api/messages"
	"github.com/WenLopes/bank-transactions-api/api/presenters"
	"github.com/WenLopes/bank-transactions-api/api/requests"
	"github.com/WenLopes/bank-transactions-api/api/responses"
	"github.com/WenLopes/bank-transactions-api/api/validators"
	"github.com/WenLopes/bank-transactions-api/app/account"
	"github.com/WenLopes/bank-transactions-api/domain"
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
			fmt.Println(err)
			responses.Error(writer, http.StatusUnprocessableEntity, err)
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
		// Logar erro aqui
		responses.Error(writer, http.StatusInternalServerError, errors.New(messages.GENERIC_ERROR))
		return
	}

	account, err := accountService.ExecuteDeposit(accountId, event.Amount)

	if err != nil {
		// Logar erro aqui, segundo parametro retornado pelo UpdateBalance
		responses.Error(writer, http.StatusInternalServerError, err)
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

	_, err := validators.ValidateWithDraw(event)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	accountId, err := strconv.Atoi(event.Origin)
	if err != nil {
		fmt.Println(err) // Logar erro aqui
		responses.Error(writer, http.StatusInternalServerError, errors.New(messages.GENERIC_ERROR))
		return
	}

	account := accountService.FindByAccountId(accountId)
	if (account == domain.Account{}) {
		responses.JSON(writer, http.StatusNotFound, 0)
		return
	}

	success, err := accountService.ExecuteWithDraw(account, event.Amount)

	if !success {
		fmt.Println(err) // Logar erro aqui
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	account = accountService.FindByAccountId(accountId)
	presenter := presenters.NewWithDrawPresenter(account.Id, account.Balance)
	responses.JSON(writer, http.StatusOK, presenter)
}

func handleTransfer(
	writer http.ResponseWriter,
	event requests.EventRequest,
	accountService account.UseCase,
) {

	_, err := validators.ValidateTransfer(event)
	if err != nil {
		responses.Error(writer, http.StatusUnprocessableEntity, err)
		return
	}

	originId, err := strconv.Atoi(event.Origin)
	if err != nil {
		fmt.Println(err) // Logar erro aqui
		responses.Error(writer, http.StatusInternalServerError, errors.New(messages.GENERIC_ERROR))
		return
	}

	destinationId, err := strconv.Atoi(event.Destination)
	if err != nil {
		fmt.Println(err) // Logar erro aqui
		responses.Error(writer, http.StatusInternalServerError, errors.New(messages.GENERIC_ERROR))
		return
	}

	accountOrigin := accountService.FindByAccountId(originId)
	if (accountOrigin == domain.Account{}) {
		fmt.Println("conta de origem não existe") // logar que conta de origem não existe
		responses.JSON(writer, http.StatusNotFound, 0)
		return
	}

	accountDestination := accountService.FindByAccountId(destinationId)
	if (accountDestination == domain.Account{}) {
		fmt.Println("conta de destino não existe") // logar que conta de destino não existe
		responses.JSON(writer, http.StatusNotFound, 0)
		return
	}

	success, err := accountService.ExecuteTransfer(accountOrigin, accountDestination, event.Amount)

	if !success {
		fmt.Println(err) // Logar erro aqui
		responses.Error(writer, http.StatusInternalServerError, err)
		return
	}

	accountOrigin = accountService.FindByAccountId(originId)
	accountDestination = accountService.FindByAccountId(destinationId)

	presenter := presenters.NewTransferPresenter(accountOrigin, accountDestination)
	responses.JSON(writer, http.StatusOK, presenter)
}

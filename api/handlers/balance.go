package handlers

import (
	"net/http"

	"github.com/WenLopes/bank-transactions-api/api/requests"
	"github.com/WenLopes/bank-transactions-api/api/responses"
	"github.com/WenLopes/bank-transactions-api/app/account"
	"github.com/WenLopes/bank-transactions-api/domain"
	"github.com/gorilla/mux"
)

func NewBalanceHandler(router *mux.Router, accountService account.UseCase) {
	router.HandleFunc("/balance", getBalance(accountService)).Methods("GET").Name("getBalance")
}

func getBalance(accountService account.UseCase) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		balanceRequest, err := requests.MapRequestToBalance(request)

		if err != nil {
			responses.Error(writer, http.StatusUnprocessableEntity, err)
			return
		}

		account := accountService.FindByAccountId(balanceRequest.AccountId)
		if (account == domain.Account{}) {
			responses.JSON(writer, http.StatusNotFound, 0)
			return
		}

		responses.JSON(writer, http.StatusOK, account.Balance)
	}
}

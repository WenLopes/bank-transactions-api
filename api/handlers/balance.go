package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/WenLopes/bank-transactions-api/api/responses"
	"github.com/WenLopes/bank-transactions-api/app/account"
	"github.com/WenLopes/bank-transactions-api/domain"
	"github.com/gorilla/mux"
)

func NewBalanceHandler(router *mux.Router, accountService account.UseCase) {
	router.HandleFunc("/balance", getBalance(accountService)).Methods("GET").Name("getBalance")
}

func getBalance(accountService account.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountIdParamExists := r.URL.Query().Has("account_id")

		if !accountIdParamExists {
			notExistsAccountIdParamError := errors.New("erro na formatação do request")
			responses.Error(w, http.StatusNotFound, notExistsAccountIdParamError)
			return
		}

		accountId, err := strconv.Atoi(r.URL.Query().Get("account_id"))

		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}

		account := accountService.FindByAccountId(accountId)
		if (account == domain.Account{}) {
			responses.JSON(w, http.StatusNotFound, 0)
			return
		}

		responses.JSON(w, http.StatusOK, account.Balance)
	}
}

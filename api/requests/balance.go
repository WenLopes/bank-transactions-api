package requests

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/WenLopes/bank-transactions-api/api/messages"
)

type BalanceRequest struct {
	AccountId int
}

func MapRequestToBalance(request *http.Request) (BalanceRequest, error) {

	accountIdParamExists := request.URL.Query().Has("account_id")

	if !accountIdParamExists {
		return BalanceRequest{}, errors.New(messages.REQUEST_INVALID)
	}

	accountId, err := strconv.Atoi(request.URL.Query().Get("account_id"))

	if err != nil {
		return BalanceRequest{}, errors.New(messages.GENERIC_ERROR)
	}

	balanceRequest := BalanceRequest{
		AccountId: accountId,
	}

	return balanceRequest, nil
}

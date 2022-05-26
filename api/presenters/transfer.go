package presenters

import (
	"strconv"

	"github.com/WenLopes/bank-transactions-api/domain"
)

type transferPresenter struct {
	Origin      idWithBalance `json:"origin"`
	Destination idWithBalance `json:"destination"`
}

func NewTransferPresenter(accountOrigin, accountDestination domain.Account) transferPresenter {
	return transferPresenter{
		Origin: idWithBalance{
			Id:      strconv.Itoa(accountOrigin.Id),
			Balance: accountOrigin.Balance,
		},
		Destination: idWithBalance{
			Id:      strconv.Itoa(accountDestination.Id),
			Balance: accountDestination.Balance,
		},
	}
}

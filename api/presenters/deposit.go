package presenters

import "strconv"

type depositPresenter struct {
	Destination idWithBalance `json:"destination"`
}

func NewDepositPresenter(id int, balance float32) depositPresenter {
	return depositPresenter{
		Destination: idWithBalance{
			Id:      strconv.Itoa(id),
			Balance: balance,
		},
	}
}

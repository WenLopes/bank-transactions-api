package presenters

import "strconv"

type withDrawPresenter struct {
	Origin idWithBalance `json:"origin"`
}

func NewWithDrawPresenter(id int, balance float32) withDrawPresenter {
	return withDrawPresenter{
		Origin: idWithBalance{
			Id:      strconv.Itoa(id),
			Balance: balance,
		},
	}
}

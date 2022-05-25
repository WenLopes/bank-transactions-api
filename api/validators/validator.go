package validators

import (
	"errors"

	"github.com/WenLopes/bank-transactions-api/api/requests"
)

var validEventTypes = [3]string{"deposit", "transfer", "withdraw"}

func ValidateDeposit(event requests.EventRequest) (bool, error) {
	valid, err := eventTypeIsValid(event.EventType)
	if !valid {
		return false, err
	}

	valid, err = amountIsValid(event.Amount)
	if !valid {
		return false, err
	}

	valid, err = destinationIsValid(event.Destination)
	if !valid {
		return false, err
	}

	return true, nil
}

func eventTypeIsValid(eventType string) (bool, error) {

	if eventType == "" {
		return false, errors.New("o tipo do evento não é válido")
	}

	for _, validEventType := range validEventTypes {
		if validEventType == eventType {
			return true, nil
		}
	}

	return false, errors.New("o tipo do evento não é válido")
}

func amountIsValid(amount float32) (bool, error) {

	if amount < 0 {
		return false, errors.New("o valor informado não é válido")
	}

	return true, nil
}

func destinationIsValid(destination string) (bool, error) {
	if destination == "" {
		return false, errors.New("o campo destination é obrigatório para realizar essa operação")
	}

	return true, nil
}

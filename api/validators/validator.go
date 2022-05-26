package validators

import (
	"errors"

	"github.com/WenLopes/bank-transactions-api/api/messages"
	"github.com/WenLopes/bank-transactions-api/api/requests"
)

var validEventTypes = [3]string{"deposit", "transfer", "withdraw"}

func ValidateDeposit(event requests.EventRequest) (bool, error) {
	valid, err := validateCommonEvent(event)
	if !valid {
		return false, err
	}

	valid, err = destinationIsValid(event.Destination)
	if !valid {
		return false, err
	}

	return true, nil
}

func ValidateTransfer(event requests.EventRequest) (bool, error) {
	valid, err := validateCommonEvent(event)
	if !valid {
		return false, err
	}

	valid, err = originIsValid(event.Origin)
	if !valid {
		return false, err
	}

	valid, err = destinationIsValid(event.Destination)
	if !valid {
		return false, err
	}

	return true, nil
}

func ValidateWithDraw(event requests.EventRequest) (bool, error) {
	valid, err := validateCommonEvent(event)
	if !valid {
		return false, err
	}

	valid, err = originIsValid(event.Origin)
	if !valid {
		return false, err
	}

	return true, nil
}

func validateCommonEvent(event requests.EventRequest) (bool, error) {

	valid, err := eventTypeIsValid(event.EventType)
	if !valid {
		return false, err
	}

	valid, err = amountIsValid(event.Amount)
	if !valid {
		return false, err
	}

	return true, nil
}

func eventTypeIsValid(eventType string) (bool, error) {

	if eventType == "" {
		return false, errors.New(messages.INVALID_EVENT_TYPE)
	}

	for _, validEventType := range validEventTypes {
		if validEventType == eventType {
			return true, nil
		}
	}

	return false, errors.New(messages.INVALID_EVENT_TYPE)
}

func amountIsValid(amount float32) (bool, error) {

	if amount < 0 {
		return false, errors.New(messages.INVALID_AMOUNT_VALUE)
	}

	return true, nil
}

func destinationIsValid(destination string) (bool, error) {
	if destination == "" {
		return false, errors.New(messages.DESTINATION_FIELD_REQUIRED)
	}

	return true, nil
}

func originIsValid(origin string) (bool, error) {
	if origin == "" {
		return false, errors.New(messages.ORIGIN_FIELD_REQUIRED)
	}

	return true, nil
}

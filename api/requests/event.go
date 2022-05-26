package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/WenLopes/bank-transactions-api/api/messages"
)

type EventRequest struct {
	EventType   string  `json:"type"`
	Origin      string  `json:"origin"`
	Amount      float32 `json:"amount"`
	Destination string  `json:"destination"`
}

func MapRequestToEvent(request *http.Request) (EventRequest, error) {
	bodyRequest, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err) // Logar erro aqui
		return EventRequest{}, errors.New(messages.GENERIC_ERROR)
	}

	var event EventRequest
	if err = json.Unmarshal(bodyRequest, &event); err != nil {
		fmt.Println(err) // Logar erro aqui
		return EventRequest{}, errors.New(messages.GENERIC_ERROR)
	}

	return event, nil
}

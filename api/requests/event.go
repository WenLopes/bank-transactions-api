package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
		return EventRequest{}, errors.New("erro ao processar requisição")
	}

	var event EventRequest
	if err = json.Unmarshal(bodyRequest, &event); err != nil {
		fmt.Println(err) // Logar erro aqui
		return EventRequest{}, errors.New("desculpe, não foi possível processar sua requisição")
	}

	return event, nil
}

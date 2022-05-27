package responses

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/WenLopes/bank-transactions-api/api/messages"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Text(w http.ResponseWriter, statusCode int, data string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte(data))
}

func Error(w http.ResponseWriter, statusCode int, err error) {

	message := messages.Get(err.Error())

	if message.ToString == "" {
		message = messages.Get(messages.GENERIC_ERROR)
	}

	JSON(w, statusCode, struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    message.Code,
		Message: message.ToString,
	})
}

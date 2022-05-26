package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WenLopes/bank-transactions-api/api/handlers"
	"github.com/WenLopes/bank-transactions-api/api/responses"
	"github.com/WenLopes/bank-transactions-api/app/account"
	"github.com/WenLopes/bank-transactions-api/domain"
	"github.com/WenLopes/bank-transactions-api/infra/repository"
	"github.com/gorilla/mux"
)

func main() {

	accountRepository := repository.NewAccountInStateRepository()
	accountService := account.NewService(accountRepository)

	router := mux.NewRouter()
	router.HandleFunc("/", home(accountRepository)).Methods("GET") //debug

	handlers.NewEventHandler(router, accountService)
	handlers.NewBalanceHandler(router, accountService)
	handlers.NewResetHandler(router, accountService)

	fmt.Printf("Escutando na porta LOCAL 16091 🏆 \n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(accountRepo domain.AccountRepository) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		responses.JSON(w, 200, accountRepo.GetAll())
	}

}

package account

import "github.com/WenLopes/bank-transactions-api/domain"

type UseCase interface {
	FindByAccountId(accountId int) domain.Account
	DeleteAll()
	ExecuteDeposit(accountId int, balance float32) (domain.Account, error)
	ExecuteWithDraw(account domain.Account, balance float32) (bool, error)
	ExecuteTransfer(accountOrigin domain.Account, accountDestinationId int, amount float32) (bool, error)
}

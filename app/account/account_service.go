package account

import (
	"github.com/WenLopes/bank-transactions-api/domain"
)

type service struct {
	accountRepo domain.AccountRepository
}

func NewService(accountRepo domain.AccountRepository) service {
	return service{
		accountRepo: accountRepo,
	}
}

func (accountService service) ExecuteDeposit(accountId int, balance float32) (domain.Account, error) {
	return domain.Account{}, nil
}

func (accountService service) ExecuteWithDraw(account domain.Account, amount float32) (bool, error) {
	return true, nil
}

func (accountService service) ExecuteTransfer(accountOrigin domain.Account, accountDestination domain.Account, amount float32) (bool, error) {
	return true, nil
}

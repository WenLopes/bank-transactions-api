package account

import (
	"errors"
	"fmt"
	"time"

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

func (accountService service) FindByAccountId(accountId int) domain.Account {
	return accountService.accountRepo.Find(accountId)
}

func (accountService service) DeleteAll() {
	accountService.accountRepo.DeleteAll()
}

func (accountService service) ExecuteDeposit(accountId int, balance float32) (domain.Account, error) {
	if balance < 0 {
		return domain.Account{}, errors.New(domain.ErrInvalidAmount)
	}

	existingAccount := accountService.accountRepo.Find(accountId)

	if (existingAccount == domain.Account{}) {
		fmt.Println("conta não existe, vai ser criada")
		account := domain.Account{
			Id:        accountId,
			Balance:   balance,
			CreatedAt: time.Now(),
		}
		accountService.accountRepo.Add(account)
		return account, nil
	}

	newBalance := (existingAccount.Balance + balance)
	success, err := accountService.accountRepo.UpdateBalance(accountId, newBalance)

	if !success {
		return domain.Account{}, err
	}

	return accountService.accountRepo.Find(accountId), nil
}

func (accountService service) ExecuteWithDraw(account domain.Account, amount float32) (bool, error) {

	if account.Locked {
		return false, errors.New(domain.BlockedAccount)
	}

	accountService.accountRepo.Lock(account.Id)
	defer accountService.accountRepo.Unlock(account.Id)

	if account.Balance < amount {
		return false, errors.New(domain.InsuficientAccountBalance)
	}

	newBalance := (account.Balance - amount)
	return accountService.accountRepo.UpdateBalance(account.Id, newBalance)
}

func (accountService service) ExecuteTransfer(accountOrigin domain.Account, accountDestination domain.Account, amount float32) (bool, error) {
	return true, nil
}

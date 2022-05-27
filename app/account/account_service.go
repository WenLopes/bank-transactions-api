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
		account := accountService.createAccount(accountId, balance)
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
		return false, errors.New(domain.InsufficientAccountBalance)
	}

	newBalance := (account.Balance - amount)
	return accountService.accountRepo.UpdateBalance(account.Id, newBalance)
}

func (accountService service) ExecuteTransfer(accountOrigin domain.Account, accountDestinationId int, amount float32) (bool, error) {

	if accountOrigin.Locked {
		return false, errors.New(domain.BlockedAccount)
	}

	accountService.accountRepo.Lock(accountOrigin.Id)
	defer accountService.accountRepo.Unlock(accountOrigin.Id)

	if accountOrigin.Balance < amount {
		return false, errors.New(domain.InsufficientAccountBalance)
	}

	accountDestination := accountService.FindByAccountId(accountDestinationId)
	if (accountDestination == domain.Account{}) {
		accountDestination = accountService.createAccount(accountDestinationId, 0)
	}

	originNewBalance := (accountOrigin.Balance - amount)
	destinationNewBalance := (accountDestination.Balance + amount)

	success, err := accountService.accountRepo.UpdateBalance(accountOrigin.Id, originNewBalance)

	if !success {
		return false, err
	}

	success, err = accountService.accountRepo.UpdateBalance(accountDestination.Id, destinationNewBalance)

	if !success {
		fmt.Println(err)

		rollback, err := accountService.rollbackTransfer(accountOrigin, accountDestination)
		if !rollback {
			fmt.Println(err)
			// Cron de conciliação
		}
		return false, err
	}

	return true, nil
}

func (accountService service) createAccount(accountId int, balance float32) domain.Account {
	account := domain.Account{
		Id:        accountId,
		Balance:   balance,
		CreatedAt: time.Now(),
	}
	accountService.accountRepo.Add(account)
	return account
}

func (accountService service) rollbackTransfer(accountOrigin, accountDestination domain.Account) (bool, error) {
	originSuccess, originErr := accountService.accountRepo.UpdateBalance(accountOrigin.Id, accountOrigin.Balance)

	if !originSuccess {
		return false, originErr
	}

	destinationSuccess, destinationErr := accountService.accountRepo.UpdateBalance(accountDestination.Id, accountDestination.Balance)

	if !destinationSuccess {
		return false, destinationErr
	}

	return true, nil
}

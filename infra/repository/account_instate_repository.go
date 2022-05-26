package repository

import "github.com/WenLopes/bank-transactions-api/domain"

var accountsState = []domain.Account{}

type accountInStateRepository struct {
	accounts []domain.Account
}

func NewAccountInStateRepository() *accountInStateRepository {
	return &accountInStateRepository{
		accounts: accountsState,
	}
}

func (repo *accountInStateRepository) Find(accountId int) domain.Account {
	account := domain.Account{}

	if len(repo.accounts) <= 0 {
		return account
	}

	for i := range repo.accounts {
		if repo.accounts[i].Id == accountId {
			account = repo.accounts[i]
		}
	}

	return account
}

func (repo *accountInStateRepository) DeleteAll() (bool, error) {
	repo.accounts = []domain.Account{}
	return true, nil
}

func (repo *accountInStateRepository) Add(account domain.Account) {
	repo.accounts = append(repo.accounts, account)
}

func (repo *accountInStateRepository) UpdateBalance(accountId int, balance float32) (bool, error) {
	accountIndex, _ := repo.getAccountIndex(accountId)
	repo.accounts[accountIndex].Balance = balance
	return true, nil
}

func (repo *accountInStateRepository) Lock(accountId int) {
	accountIndex, _ := repo.getAccountIndex(accountId)
	repo.accounts[accountIndex].Locked = true
}

func (repo *accountInStateRepository) Unlock(accountId int) {
	accountIndex, _ := repo.getAccountIndex(accountId)
	repo.accounts[accountIndex].Locked = false
}

func (repo *accountInStateRepository) getAccountIndex(accountId int) (int, bool) {
	idx := 0
	exist := false

	if len(repo.accounts) <= 0 {
		return idx, exist
	}

	for i := range repo.accounts {
		if repo.accounts[i].Id == accountId {
			idx = i
			exist = true
		}
	}

	return idx, exist
}

func (repo *accountInStateRepository) GetAll() []domain.Account {
	return repo.accounts
}

package domain

import "time"

type Account struct {
	Id        int
	Balance   float32
	Locked    bool
	CreatedAt time.Time
}

type AccountRepository interface {
	Find(accountId int) Account
	DeleteAll() (bool, error)
	Add(account Account)
	UpdateBalance(accountId int, balance float32) (bool, error)
	Lock(accountId int)
	Unlock(accountId int)
	GetAll() []Account
}

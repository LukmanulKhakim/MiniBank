package domain

import "time"

type AccountCore struct {
	Account_no uint
	Name       string
	Balance    float64
	Created_at time.Time
	Updated_at time.Time
}

type Repository interface {
	Add(new AccountCore) (AccountCore, error)
	GetAll(input bool) ([]AccountCore, error)
	GetMyUser(userID uint) (AccountCore, error)
}

type Service interface {
	AddAccount(new AccountCore) (AccountCore, error)
	AllAccount(input bool) ([]AccountCore, error)
	MyAccount(userID uint) (AccountCore, error)
}

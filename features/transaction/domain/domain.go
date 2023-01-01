package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionCore struct {
	Transaction_id uuid.UUID
	Credit_account uint
	Debit_account  uint
	Amount         float64
	Created_at     time.Time
}

type AccountCore struct {
	Account_no uint
	Name       string
	Balance    float64
	Created_at time.Time
	Updated_at time.Time
}

type Repository interface {
	Add(new TransactionCore) (TransactionCore, error)
	GetAll() ([]TransactionCore, error)
	Get() ([]TransactionCore, error)
	Filter(filter []TransactionCore) ([]TransactionCore, error)
}

type Service interface {
	AddTrx(new TransactionCore) (TransactionCore, error)
	AllTrx() ([]TransactionCore, error)
	Trx() ([]TransactionCore, error)
	Search(filter []TransactionCore) ([]TransactionCore, error)
}

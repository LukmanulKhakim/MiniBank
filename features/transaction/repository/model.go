package repository

import (
	"minibank/features/transaction/domain"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Balance     float64
	Transaction []Transaction `gorm:"foreignKey:Credit_account;foreignKey:Debit_account"`
}

type Base struct {
	ID        uuid.UUID `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}

type Transaction struct {
	Base
	Credit_account uint
	Debit_account  uint
	Amount         float64
}

func FromDomain(du domain.TransactionCore) Transaction {
	return Transaction{
		Base:           Base{ID: du.Transaction_id, CreatedAt: du.Created_at},
		Credit_account: du.Credit_account,
		Debit_account:  du.Debit_account,
		Amount:         du.Amount,
	}
}

func ToDomain(u Transaction) domain.TransactionCore {
	return domain.TransactionCore{
		Transaction_id: u.ID,
		Credit_account: u.Credit_account,
		Debit_account:  u.Debit_account,
		Amount:         u.Amount,
		Created_at:     u.CreatedAt,
	}
}

func ToDomainArray(av []Transaction) []domain.TransactionCore {
	var res []domain.TransactionCore
	for _, val := range av {
		res = append(res, domain.TransactionCore{
			Transaction_id: val.ID,
			Credit_account: val.Credit_account,
			Debit_account:  val.Debit_account,
			Amount:         val.Amount,
			Created_at:     val.CreatedAt,
		})
	}
	return res
}

func FromDomainA(du domain.AccountCore) Account {
	return Account{
		Model:   gorm.Model{ID: du.Account_no, CreatedAt: du.Created_at},
		Name:    du.Name,
		Balance: du.Balance,
	}
}

func ToDomainA(a Account) domain.AccountCore {
	return domain.AccountCore{
		Account_no: a.ID,
		Name:       a.Name,
		Balance:    a.Balance,
		Created_at: a.CreatedAt,
		Updated_at: a.UpdatedAt,
	}
}

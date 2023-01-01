package repository

import (
	"minibank/features/account/domain"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Balance float64
}

func FromDomain(du domain.AccountCore) Account {
	return Account{
		Model:   gorm.Model{ID: du.Account_no, CreatedAt: du.Created_at},
		Name:    du.Name,
		Balance: du.Balance,
	}
}

func ToDomain(u Account) domain.AccountCore {
	return domain.AccountCore{
		Account_no: u.ID,
		Name:       u.Name,
		Balance:    u.Balance,
		Created_at: u.CreatedAt,
	}
}

func ToDomainArray(av []Account) []domain.AccountCore {
	var res []domain.AccountCore
	for _, val := range av {
		res = append(res, domain.AccountCore{
			Account_no: val.ID,
			Name:       val.Name,
			Balance:    val.Balance,
			Created_at: val.CreatedAt,
		})
	}
	return res
}

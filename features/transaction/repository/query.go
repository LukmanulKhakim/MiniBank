package repository

import (
	"fmt"
	"minibank/features/transaction/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var db *gorm.DB

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{
		db: db,
	}
}

// Add implements domain.Repository
func (rq *repoQuery) Add(new domain.TransactionCore) (domain.TransactionCore, error) {
	var cnv Transaction = FromDomain(new)
	var cnv2, cnv3 Account
	var res2, res3 float64

	row := rq.db.Table("accounts").Where("id", new.Debit_account).Select("balance").Row()
	row.Scan(&res2)

	row2 := rq.db.Table("accounts").Where("id", new.Credit_account).Select("balance").Row()
	row2.Scan(&res3)

	cnv.CreatedAt = new.Created_at
	cnv2.Balance = res2 + cnv.Amount
	cnv3.Balance = res3 - cnv.Amount
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on adding user", err.Error())
		return domain.TransactionCore{}, err
	}
	if err := rq.db.Table("accounts").Where("id = ?", new.Debit_account).Updates(&cnv2).Error; err != nil {
		log.Error("query get Account", err.Error())
		return domain.TransactionCore{}, err
	}

	if err := rq.db.Table("accounts").Where("id = ?", new.Credit_account).Updates(&cnv3).Error; err != nil {
		log.Error("query get Account", err.Error())
		return domain.TransactionCore{}, err
	}
	return ToDomain(cnv), nil
}

// Get implements domain.Repository
func (rq *repoQuery) Get() ([]domain.TransactionCore, error) {
	var res []Transaction
	if err := rq.db.Table("transactions").Select("amount", "created_at").Find(&res).Error; err != nil {
		log.Error("error on getuserid", err.Error())
		return []domain.TransactionCore{}, err
	}
	return ToDomainArray(res), nil
}

// GetAll implements domain.Repository
func (rq *repoQuery) GetAll() ([]domain.TransactionCore, error) {
	var res []Transaction
	if err := rq.db.Find(&res).Error; err != nil {
		log.Error("error on get all user", err.Error())
		return []domain.TransactionCore{}, err
	}
	return ToDomainArray(res), nil
}

// Filter implements domain.Repository
func (rq *repoQuery) Filter(filter []domain.TransactionCore) ([]domain.TransactionCore, error) {
	var res []Transaction
	for i := 0; i < len(filter); i++ {
		if err := rq.db.Table("transactions").Where("credit_account= ? OR debit_account = ? OR amount = ?", filter[0].Credit_account, filter[1].Debit_account, filter[2].Amount).Find(&res).Error; err != nil {
			log.Error("error on get all user", err.Error())
			fmt.Println(filter)
			return []domain.TransactionCore{}, err
		}
	}
	fmt.Println(res)
	return ToDomainArray(res), nil
}

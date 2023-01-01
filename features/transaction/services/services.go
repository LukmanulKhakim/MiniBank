package services

import (
	"errors"
	"fmt"
	"minibank/config"
	"minibank/features/transaction/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type trxService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &trxService{
		qry: repo,
	}
}

// AddTrx implements domain.Service
func (ts *trxService) AddTrx(new domain.TransactionCore) (domain.TransactionCore, error) {
	res, err := ts.qry.Add(new)
	if err != nil {
		return domain.TransactionCore{}, errors.New("some problem on database")
	}
	return res, nil
}

// AllTrx implements domain.Service
func (ts *trxService) AllTrx() ([]domain.TransactionCore, error) {
	res, err := ts.qry.GetAll()
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}
	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New(config.DATA_NOTFOUND)
	}
	return res, nil
}

// Trx implements domain.Service
func (ts *trxService) Trx() ([]domain.TransactionCore, error) {
	res, err := ts.qry.Get()
	if err == gorm.ErrRecordNotFound {
		log.Error("error on service", err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error("error on service", err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}
	return res, nil
}

// Search implements domain.Service
func (ts *trxService) Search(filter []domain.TransactionCore) ([]domain.TransactionCore, error) {
	res, err := ts.qry.Filter(filter)
	if err == gorm.ErrRecordNotFound {
		fmt.Println(res)
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}
	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New(config.DATA_NOTFOUND)
	}
	return res, nil
}

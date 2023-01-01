package services

import (
	"errors"
	"minibank/config"
	"minibank/features/account/domain"
	"strings"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type accountService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &accountService{
		qry: repo,
	}
}

// AddAccount implements domain.Service
func (as *accountService) AddAccount(new domain.AccountCore) (domain.AccountCore, error) {
	res, err := as.qry.Add(new)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return domain.AccountCore{}, errors.New("already exist account")
		}
		return domain.AccountCore{}, errors.New("some problem on database")
	}
	return res, nil
}

// AllAccount implements domain.Service
func (as *accountService) AllAccount(input bool) ([]domain.AccountCore, error) {
	res, err := as.qry.GetAll(input)
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

// MyAccount implements domain.Service
func (as *accountService) MyAccount(userID uint) (domain.AccountCore, error) {
	res, err := as.qry.GetMyUser(userID)
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return domain.AccountCore{}, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return domain.AccountCore{}, errors.New(config.DATABASE_ERROR)
	}
	return res, nil
}

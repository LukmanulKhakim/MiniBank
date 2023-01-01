package repository

import (
	"minibank/features/account/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

// Add implements domain.Repository
func New(db *gorm.DB) domain.Repository {
	return &repoQuery{
		db: db,
	}
}
func (rq *repoQuery) Add(new domain.AccountCore) (domain.AccountCore, error) {
	var cnv Account = FromDomain(new)
	cnv.CreatedAt = new.Created_at
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on adding user", err.Error())
		return domain.AccountCore{}, err
	}
	return ToDomain(cnv), nil
}

// GetAll implements domain.Repository
func (rq *repoQuery) GetAll(input bool) ([]domain.AccountCore, error) {
	var res []Account
	if err := rq.db.Find(&res).Error; err != nil {
		log.Error("error on get all user", err.Error())
		return []domain.AccountCore{}, err
	}

	return ToDomainArray(res), nil
}

// GetMyUser implements domain.Repository
func (rq *repoQuery) GetMyUser(userID uint) (domain.AccountCore, error) {
	var res Account
	if err := rq.db.First(&res, "id =?", userID).Error; err != nil {
		log.Error("error on getuserid", err.Error())
		return domain.AccountCore{}, err
	}
	return ToDomain(res), nil
}

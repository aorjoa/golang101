package model

import "github.com/jinzhu/gorm"

type UserLoginLog struct {
	gorm.Model
	Username string
	Password string
}

type UserLoginLogDB interface {
	List() ([]UserLoginLog, error)
	Create(uli *UserLoginLog) error
}

type gormUserLoginLogDB struct {
	db *gorm.DB
}

func NewUserLoginLogDB(db *gorm.DB) UserLoginLogDB {
	return &gormUserLoginLogDB{
		db: db,
	}
}

func (ul *gormUserLoginLogDB) List() ([]UserLoginLog, error) {
	return nil, nil
}

func (ul *gormUserLoginLogDB) Create(uli *UserLoginLog) error {
	return ul.db.Create(uli).Error
}

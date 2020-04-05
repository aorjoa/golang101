package mock

import (
	"ssm/model"

	"github.com/stretchr/testify/mock"
)

// UserLoginLogDB is mock service
type UserLoginLogDB struct {
	mock.Mock
}

var _ model.UserLoginLogDB = &UserLoginLogDB{}

// List is a mock return
func (ul UserLoginLogDB) List() ([]model.UserLoginLog, error) {
	args := ul.Called()
	return args.Get(0).([]model.UserLoginLog), args.Error(1)
}

// Create is a mock return
func (ul UserLoginLogDB) Create(ulm *model.UserLoginLog) error {
	args := ul.Called(ulm)
	return args.Error(0)
}

package user

import (
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"go.uber.org/zap"
)

type User interface {
	CreateUser(req *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
}

type impl struct {
	logger *zap.Logger
	db     *database.SqlDB
}

func New(logger *zap.Logger, db *database.SqlDB) User {
	return &impl{
		logger: logger,
		db:     db,
	}
}

func (i *impl) CreateUser(req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, nil
}

func (i *impl) GetUser(req *GetUserRequest) (*GetUserResponse, error) {
	return nil, nil
}

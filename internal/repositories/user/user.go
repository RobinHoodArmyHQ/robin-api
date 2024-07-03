package user

import (
	"fmt"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	err := i.db.Master().Create(req.User).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &CreateUserResponse{UserID: req.User.UserId}, nil
}

func (i *impl) GetUser(req *GetUserRequest) (*GetUserResponse, error) {
	user := &models.User{}
	exec := i.db.Master().First(user, "user_id = ?", req.UserID)
	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		i.logger.Error("user not found", zap.Uint64("user_id", req.UserID))
		return nil, nil
	}

	return &GetUserResponse{
		User: user,
	}, nil
}

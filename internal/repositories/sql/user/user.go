package user

import (
	"fmt"

	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/user"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type impl struct {
	logger *zap.Logger
	db     *database.SqlDB
}

func New(logger *zap.Logger, db *database.SqlDB) user.User {
	return &impl{
		logger: logger,
		db:     db,
	}
}

func (i *impl) CreateUser(req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	var err error
	req.User.ID = 0
	req.User.UserID, err = nanoid.GetID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate user id: %v", err)
	}

	err = i.db.Master().Create(req.User).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &user.CreateUserResponse{UserID: req.User.UserID}, nil
}

func (i *impl) GetUser(req *user.GetUserRequest) (*user.GetUserResponse, error) {
	model := &models.User{}
	exec := i.db.Master().First(model, "user_id = ?", req.UserID)
	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		i.logger.Error("user not found", zap.String("user_id", req.UserID.String()))
		return nil, nil
	}

	return &user.GetUserResponse{
		User: model,
	}, nil
}

func (i *impl) CheckIfUserExists(req *user.CheckIfUserExistsRequest) (*user.CheckIfUserExistsResponse, error) {
	userData := &models.User{}
	exec := i.db.Master().First(userData, "email_id = ?", req.EmailId)

	if exec.Error != nil {
		i.logger.Error("ERROR_CHECK_IF_USER_EXISTS", zap.Error(exec.Error))
		return nil, exec.Error
	}

	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		i.logger.Info("user not found", zap.String("email_id", req.EmailId))
		return &user.CheckIfUserExistsResponse{
			IsExisting: false,
		}, nil
	}

	return &user.CheckIfUserExistsResponse{
		IsExisting: true,
	}, nil
}

func (i *impl) GetUserByEmailId(req *user.GetUserByEmailIdRequest) (*user.GetUserByEmailIdResponse, error) {
	userData := &models.User{}
	exec := i.db.Master().First(userData, "email_id=?", req.EmailId)

	if exec.Error != nil {
		i.logger.Error("ERROR_GET_USER_BY_EMAIL_ID", zap.Error(exec.Error))
		return nil, exec.Error
	}

	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		i.logger.Info("user not found", zap.String("email_id", req.EmailId))
		return nil, nil
	}

	return &user.GetUserByEmailIdResponse{
		User: userData,
	}, nil
}

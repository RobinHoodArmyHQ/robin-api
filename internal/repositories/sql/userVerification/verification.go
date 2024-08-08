package userverification

import (
	"errors"
	"fmt"

	userverification "github.com/RobinHoodArmyHQ/robin-api/internal/repositories/userVerification"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type impl struct {
	logger *zap.Logger
	db     *database.SqlDB
}

func New(logger *zap.Logger, db *database.SqlDB) userverification.UserVerification {
	return &impl{
		logger: logger,
		db:     db,
	}
}

func (i impl) CreateUser(req *userverification.CreateUserRequest) (*userverification.CreateUserResponse, error) {
	var err error
	req.User.UserID, err = nanoid.GetID()

	if err != nil {
		return nil, fmt.Errorf("failed to generate user id: %v", err)
	}

	err = i.db.Master().Create(req.User).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &userverification.CreateUserResponse{UserID: req.User.UserID}, nil
}

func (i impl) GetUserByUserID(req *userverification.GetUserByUserIdRequest) (*userverification.GetUserByUserIdResponse, error) {
	userData := &models.UserVerification{}
	exec := i.db.Master().First(userData, "user_id=?", req.UserID)

	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		i.logger.Error("user not found", zap.String("user_id", req.UserID.String()))
		return nil, nil
	}

	if exec.Error != nil {
		i.logger.Error("ERROR_GET_UNVERIFIED_USER_BY_USER_ID", zap.Error(exec.Error))
		return nil, exec.Error
	}

	return &userverification.GetUserByUserIdResponse{User: userData}, nil
}

func (i impl) UpdateUser(req *userverification.UpdateUserRequest) (*userverification.UpdateUserResponse, error) {
	model := []*models.UserVerification{}
	exec := i.db.Master().Model(model).Clauses(clause.Returning{}).Where("user_id=?", req.UserID).Updates(req.Values)

	if exec.Error != nil {
		i.logger.Error("ERROR_UPDATE_USER_VERIFICATIONS", zap.Error(exec.Error))
		return nil, exec.Error
	}

	records := userverification.UpdateUserResponse{
		Users: model,
	}

	return &records, nil
}

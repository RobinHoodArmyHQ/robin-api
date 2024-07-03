package checkin

import (
	"fmt"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CheckIn interface {
	CreateCheckIn(req *CreateCheckInRequest) (*CreateCheckInResponse, error)
	GetCheckIn(req *GetCheckInRequest) (*GetCheckInResponse, error)
	GetUserCheckIns(req *GetUserCheckInsRequest) (*GetUserCheckInsResponse, error)
}

type impl struct {
	logger *zap.Logger
	db     *database.SqlDB
}

func New(logger *zap.Logger, db *database.SqlDB) CheckIn {
	return &impl{
		logger: logger,
		db:     db,
	}
}

func (i *impl) CreateCheckIn(req *CreateCheckInRequest) (*CreateCheckInResponse, error) {
	err := i.db.Master().Create(req.CheckIn).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create checkin: %v", err)
	}

	return &CreateCheckInResponse{CheckInID: req.CheckIn.CheckInID}, nil
}

func (i *impl) GetCheckIn(req *GetCheckInRequest) (*GetCheckInResponse, error) {
	checkin := &models.CheckIn{}
	exec := i.db.Master().First(checkin, "check_in_id = ?", req.CheckInID)
	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		i.logger.Error("user not found", zap.Uint64("check_in_id", req.CheckInID))
		return nil, nil
	}

	return &GetCheckInResponse{
		CheckIn: checkin,
	}, nil
}

func (i *impl) GetUserCheckIns(req *GetUserCheckInsRequest) (*GetUserCheckInsResponse, error) {
	checkins := make([]*models.CheckIn, 0)
	exec := i.db.Master().Find(&checkins, "user_id = ?", req.UserID)
	if exec.Error != nil {
		return nil, fmt.Errorf("failed to get checkins: %v", exec.Error)
	}

	return &GetUserCheckInsResponse{
		CheckIns: checkins,
	}, nil
}

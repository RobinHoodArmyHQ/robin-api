package checkin

import (
	"fmt"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/checkin"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type impl struct {
	logger *zap.Logger
	db     *database.SqlDB
}

func New(logger *zap.Logger, db *database.SqlDB) checkin.CheckIn {
	return &impl{
		logger: logger,
		db:     db,
	}
}

func (i *impl) CreateCheckIn(req *checkin.CreateCheckInRequest) (*checkin.CreateCheckInResponse, error) {
	var err error
	req.CheckIn.ID = 0
	req.CheckIn.CheckInID, err = uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate checkin id: %v", err)
	}

	err = i.db.Master().Create(req.CheckIn).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create checkin: %v", err)
	}

	return &checkin.CreateCheckInResponse{CheckInID: req.CheckIn.CheckInID}, nil
}

func (i *impl) GetCheckIn(req *checkin.GetCheckInRequest) (*checkin.GetCheckInResponse, error) {
	model := &models.CheckIn{}
	exec := i.db.Master().First(model, "check_in_id = ?", req.CheckInID)
	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		i.logger.Error("user not found", zap.String("check_in_id", req.CheckInID.String()))
		return nil, nil
	}

	return &checkin.GetCheckInResponse{
		CheckIn: model,
	}, nil
}

func (i *impl) GetUserCheckIns(req *checkin.GetUserCheckInsRequest) (*checkin.GetUserCheckInsResponse, error) {
	checkins := make([]*models.CheckIn, 0)
	exec := i.db.Master().Find(&checkins, "user_id = ?", req.UserID)
	if exec.Error != nil {
		return nil, fmt.Errorf("failed to get checkins: %v", exec.Error)
	}

	return &checkin.GetUserCheckInsResponse{
		CheckIns: checkins,
	}, nil
}

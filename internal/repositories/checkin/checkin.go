package checkin

import (
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"go.uber.org/zap"
)

type CheckIn interface {
	CreateCheckIn(req *CreateCheckInRequest) (*CreateCheckInResponse, error)
	GetCheckIn(req *GetCheckInRequest) (*GetCheckInResponse, error)
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
	return nil, nil
}

func (i *impl) GetCheckIn(req *GetCheckInRequest) (*GetCheckInResponse, error) {
	return nil, nil
}

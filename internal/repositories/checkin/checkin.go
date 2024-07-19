package checkin

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type CheckIn interface {
	CreateCheckIn(req *CreateCheckInRequest) (*CreateCheckInResponse, error)
	GetCheckIn(req *GetCheckInRequest) (*GetCheckInResponse, error)
	GetUserCheckIns(req *GetUserCheckInsRequest) (*GetUserCheckInsResponse, error)
}

type CreateCheckInRequest struct {
	CheckIn *models.CheckIn `json:"check_in"`
}

type CreateCheckInResponse struct {
	CheckInID nanoid.NanoID `json:"check_in_id"`
}

type GetCheckInRequest struct {
	CheckInID nanoid.NanoID `json:"check_in_id"`
}

type GetCheckInResponse struct {
	CheckIn *models.CheckIn `json:"check_in"`
}

type GetUserCheckInsRequest struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type GetUserCheckInsResponse struct {
	CheckIns []*models.CheckIn `json:"check_ins"`
}

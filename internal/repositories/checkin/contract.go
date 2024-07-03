package checkin

import "github.com/RobinHoodArmyHQ/robin-api/models"

type CreateCheckInRequest struct {
	CheckIn *models.CheckIn `json:"check_in"`
}

type CreateCheckInResponse struct {
	CheckInID uint64 `json:"check_in_id"`
}

type GetCheckInRequest struct {
	CheckInID uint64 `json:"check_in_id"`
}

type GetCheckInResponse struct {
	CheckIn *models.CheckIn `json:"check_in"`
}

type GetUserCheckInsRequest struct {
	UserID uint64 `json:"user_id"`
}

type GetUserCheckInsResponse struct {
	CheckIns []*models.CheckIn `json:"check_ins"`
}

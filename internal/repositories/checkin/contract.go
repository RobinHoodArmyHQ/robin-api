package checkin

import "github.com/RobinHoodArmyHQ/robin-api/models"

type CreateCheckInRequest struct {
	CheckIn *models.CheckIn
}

type CreateCheckInResponse struct {
	CheckInID int64
}

type GetCheckInRequest struct {
	CheckInID int64
}

type GetCheckInResponse struct {
	CheckIn *models.CheckIn
}

package event

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
)

type GetEventFeedRequest struct {
	Page   int
	Limit  int
	CityId int32
}

type GetEventFeedResponse struct {
	Events []*models.Event
}

type GetEventParticipantsResponse struct {
	Participants []*models.Participant
}

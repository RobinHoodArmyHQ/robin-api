package event

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type GetEventsRequest struct {
	Page   int           `form:"page"`
	Limit  int           `form:"limit"`
	CityId nanoid.NanoID `form:"city_id"`
	// TODO: use lat, long and send nearest events on top
	//Latitude      float64 `json:"latitude"`
	//Longitude     float64 `json:"longitude"`
	//GooglePlaceID string  `json:"google_place_id,omitempty"`
}

type GetEventsResponse struct {
	Page   int             `json:"page"`
	Limit  int             `json:"limit"`
	Status *models.Status  `json:"status,omitempty"`
	Events []*models.Event `json:"events,omitempty"`
}

type InterestedEventRequest struct {
	EventID nanoid.NanoID `json:"event_id" binding:"required"`
}

type InterestedEventResponse struct {
	Status *models.Status `json:"status,omitempty"`
}

type GetParticipantsResponse struct {
	Status       *models.Status        `json:"status,omitempty"`
	Participants []*models.Participant `json:"participants,omitempty"`
}

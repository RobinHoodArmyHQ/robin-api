package event

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type GetEventsRequest struct {
	Offset       int             `form:"offset"`
	Limit        int             `form:"limit"`
	CityId       nanoid.NanoID   `form:"city_id"`
	UserLocation models.Location `form:"user_location"`
	// TODO: use lat, long and send nearest events on top
	//Latitude      float64 `json:"latitude"`
	//Longitude     float64 `json:"longitude"`
	//GooglePlaceID string  `json:"google_place_id,omitempty"`
}

type GetEventsResponse struct {
	Status *models.Status  `json:"status"`
	Events []*models.Event `json:"events"`
	Offset int             `json:"offset"`
	Count  int             `json:"count"`
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

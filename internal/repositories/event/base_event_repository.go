package event

import "github.com/RobinHoodArmyHQ/robin-api/models"

type EventRepository interface {
	CreateEvent(event *CreateEventRequest) (*CreateEventResponse, error)
}

type CreateEventRequest struct {
	Event *models.Event
}

type CreateEventResponse struct {
	EventID int64
}

package event

import "github.com/RobinHoodArmyHQ/robin-api/models"

type EventRepository interface {
	CreateEvent(req *CreateEventRequest) (*CreateEventResponse, error)
	GetEvent(req *GetEventRequest) (*GetEventResponse, error)
}

type CreateEventRequest struct {
	Event *models.Event
}

type CreateEventResponse struct {
	EventID string
}

type GetEventRequest struct {
	EventID string
}

type GetEventResponse struct {
	Event *models.Event
}

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
	EventID int64
}

type GetEventRequest struct {
	EventID int64
}

type GetEventResponse struct {
	Event *models.Event
}

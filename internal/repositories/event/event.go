package event

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type EventRepository interface {
	CreateEvent(req *CreateEventRequest) (*CreateEventResponse, error)
	GetEvent(req *GetEventRequest) (*GetEventResponse, error)
	GetEventsByLocation(req *GetEventsByLocationRequest) (*GetEventsByLocationResponse, error)
}

type CreateEventRequest struct {
	Event *models.Event
}

type CreateEventResponse struct {
	EventID nanoid.NanoID
}

type GetEventRequest struct {
	EventID nanoid.NanoID
}

type GetEventResponse struct {
	Event *models.Event
}

type GetEventsByLocationRequest struct {
	CityId int32
	Limit  int
	Offset int
}

type GetEventsByLocationResponse struct {
	Events []*models.Event
}

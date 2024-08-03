package event

import (
	"errors"
	"fmt"
	eventModel "github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	PreloadEventLocation = "EventLocation"
)

// EventRepository is the repository for events
type EventRepository struct {
	logger *zap.Logger
	db     *database.SqlDB
}

// NewEventRepository creates a new EventRepository
func NewEventRepository(logger *zap.Logger, db *database.SqlDB) *EventRepository {
	return &EventRepository{logger, db}
}

// CreateEvent creates a new event within a transaction
func (r *EventRepository) CreateEvent(req *eventModel.CreateEventRequest) (*eventModel.CreateEventResponse, error) {
	id, err := nanoid.GetID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate nanoid: %v", err)
	}
	req.Event.EventId = id

	err = r.db.Master().Create(req.Event).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create event: %v", err)
	}

	return &eventModel.CreateEventResponse{EventID: req.Event.EventId}, nil
}

// GetEvent retrieves an event by its ID
func (r *EventRepository) GetEvent(req *eventModel.GetEventRequest) (*eventModel.GetEventResponse, error) {
	event := &models.Event{}
	exec := r.db.Master().
		Preload(PreloadEventLocation).
		Preload("EventLocation.City").
		Preload("EventLocation.City.Country").
		First(event, "event_id = ?", req.EventID)

	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		err := fmt.Errorf("event not found: %v", req.EventID)
		r.logger.Error(err.Error())
		return &eventModel.GetEventResponse{}, err
	}

	return &eventModel.GetEventResponse{
		Event: event,
	}, nil
}

func (r *EventRepository) GetEventsByLocation(req *eventModel.GetEventsByLocationRequest) (*eventModel.GetEventsByLocationResponse, error) {
	var events []*models.Event
	//now := time.Now().Format(timeFormat)
	exec := r.db.Master().
		Preload(PreloadEventLocation).
		Preload("EventLocation.City").
		Preload("EventLocation.City.Country").
		Where("location_id IN (SELECT id FROM locations WHERE city_id = ?)", req.CityId).
		Order(`
			CASE 
				WHEN start_time >= NOW() THEN 0
				ELSE 1
			END,
			CASE 
				WHEN start_time >= NOW() THEN TIMESTAMPDIFF(SECOND, start_time, NOW())
				ELSE TIMESTAMPDIFF(SECOND, NOW(), start_time)
			END DESC
		`).
		Offset(req.Offset).
		Limit(req.Limit).
		Find(&events)
	if exec.Error != nil {
		return nil, fmt.Errorf("failed to get events: %v", exec.Error)
	}

	return &eventModel.GetEventsByLocationResponse{
		Events: events,
	}, nil
}

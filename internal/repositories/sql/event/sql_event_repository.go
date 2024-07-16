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
	req.Event.EventId = id.String()

	err = r.db.Master().Create(req.Event).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create event: %v", err)
	}

	return &eventModel.CreateEventResponse{EventID: req.Event.EventId}, nil
}

// GetEvent retrieves an event by its ID
func (r *EventRepository) GetEvent(req *eventModel.GetEventRequest) (*eventModel.GetEventResponse, error) {
	event := &models.Event{}
	exec := r.db.Master().Preload(PreloadEventLocation).First(event, "event_id = ?", req.EventID)
	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		r.logger.Error("event not found", zap.String("event_id", req.EventID))
		return &eventModel.GetEventResponse{}, nil
	}

	return &eventModel.GetEventResponse{
		Event: event,
	}, nil
}

package sql

import (
	"database/sql"
	"fmt"

	eventModel "github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
)

// EventRepository is the repository for events
type EventRepository struct {
	db *sql.DB
}

// NewEventRepository creates a new EventRepository
func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db}
}

// CreateEvent creates a new event
func (r *EventRepository) CreateEvent(req *eventModel.CreateEventRequest) (*eventModel.CreateEventResponse, error) {
	event := req.Event
	query := `INSERT INTO events (name, description) VALUES (?, ?)`
	res, err := r.db.Exec(query, event.Name, event.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to create event: %v", err)
	}

	eventID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get event ID: %v", err)
	}

	return &eventModel.CreateEventResponse{EventID: eventID}, nil
}

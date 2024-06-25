package sql

import (
	"database/sql"
	"fmt"
	"time"

	eventModel "github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
	"github.com/RobinHoodArmyHQ/robin-api/models"
)

// EventRepository is the repository for events
type EventRepository struct {
	db *sql.DB
}

// NewEventRepository creates a new EventRepository
func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db}
}

// CreateEvent creates a new event within a transaction
func (r *EventRepository) CreateEvent(req *eventModel.CreateEventRequest) (*eventModel.CreateEventResponse, error) {
	event := req.Event
	var locationID *int64

	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Ensure to rollback the transaction in case of an error or panic
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw the panic after rollback
		} else if err != nil {
			tx.Rollback()
		}
	}()

	if event.EventLocation != nil {
		locQuery := `INSERT INTO locations (name, latitude, longitude) VALUES (?, ?, ?)`
		locRes, err := tx.Exec(locQuery, event.EventLocation.Name, event.EventLocation.Latitude, event.EventLocation.Longitude)
		if err != nil {
			return nil, fmt.Errorf("failed to create event location: %v", err)
		}

		id, err := locRes.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("failed to get location ID: %v", err)
		}
		locationID = &id
	}

	query := `INSERT INTO events (name, description, start_time, event_type, event_location_id, min_robins, max_robins, created_at, created_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := tx.Exec(query, event.Name, event.Description, event.StartTime, event.EventType, locationID, event.MinRobins, event.MaxRobins, time.Now(), event.CreatedBy)
	if err != nil {
		return nil, fmt.Errorf("failed to create event: %v", err)
	}

	eventID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get event ID: %v", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &eventModel.CreateEventResponse{EventID: eventID}, nil
}

// GetEvent retrieves an event by its ID
func (r *EventRepository) GetEvent(req *eventModel.GetEventRequest) (*eventModel.GetEventResponse, error) {
	query := `
		SELECT e.event_id, e.name, e.description, e.start_time, e.event_type, e.event_location_id, e.min_robins, e.max_robins, e.created_at, e.created_by,
		e.updated_at, e.updated_by, l.name, l.latitude, l.longitude
		FROM events e
		LEFT JOIN locations l ON e.event_location_id = l.location_id
		WHERE e.event_id = ?`

	row := r.db.QueryRow(query, req.EventID)

	var event models.Event
	var location models.Location

	err := row.Scan(&event.EventId, &event.Name, &event.Description, &event.StartTime, &event.EventType, &event.EventLocationID,
		&event.MinRobins, &event.MaxRobins, &event.CreatedAt, &event.CreatedBy, &event.UpdatedAt, &event.UpdatedBy,
		&location.Name, &location.Latitude, &location.Longitude)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no event found with ID %d", req.EventID)
		}
		return nil, fmt.Errorf("failed to retrieve event: %s", err.Error())
	}

	if event.EventLocationID != 0 {
		event.EventLocation = &location
	}

	return &eventModel.GetEventResponse{
		Event: &event,
	}, nil
}

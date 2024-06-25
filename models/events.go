package models

import (
	"database/sql"
	"time"
)

type EventType uint8

const (
	EventInvalid EventType = iota
	EventMealDrive
	EventAcademy
)

type VehicleType uint8

const (
	VehicleInvalid VehicleType = iota
	VehicleTwoWheeler
	VehicleFourWheeler
)

var (
	ValidEventTypes = map[EventType]bool{
		EventMealDrive: true,
		EventAcademy:   true,
	}

	ValidVehicleTypes = map[VehicleType]bool{
		VehicleTwoWheeler:  true,
		VehicleFourWheeler: true,
	}
)

type Event struct {
	EventId         int64         `json:"event_id,omitempty"`
	Name            string        `json:"name,omitempty"`
	Description     string        `json:"description,omitempty"`
	StartTime       time.Time     `json:"start_time,omitempty"`
	EventType       EventType     `json:"event_type,omitempty"`
	EventLocationID int64         `json:"-"`
	EventLocation   *Location     `json:"event_location,omitempty"`
	MinRobins       uint8         `json:"min_robins,omitempty"`
	MaxRobins       uint8         `json:"max_robins,omitempty"`
	CreatedBy       int64         `json:"-"`
	UpdatedBy       sql.NullInt64 `json:"-"`
	CreatedAt       time.Time     `json:"created_at,omitempty"`
	UpdatedAt       sql.NullTime  `json:"updated_at,omitempty"`
}

type CreateEventResponse struct {
	Status  *Status `json:"status,omitempty"`
	EventId int64   `json:"event_id,omitempty"`
}

type GetEventResponse struct {
	Status *Status `json:"status,omitempty"`
	Event  *Event  `json:"event,omitempty"`
}

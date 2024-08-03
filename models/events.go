package models

import (
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type EventType string

const (
	EventMealDrive EventType = "MEAL_DRIVE"
	EventAcademy   EventType = "ACADEMY"
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
	ID            int64         `json:"-" gorm:"primaryKey"`
	EventId       nanoid.NanoID `json:"event_id,omitempty"`
	Title         string        `json:"title,omitempty"`
	Description   string        `json:"description,omitempty"`
	StartTime     time.Time     `json:"start_time,omitempty"`
	EventType     EventType     `json:"event_type,omitempty"`
	LocationID    int64         `json:"-"`
	EventLocation *Location     `json:"event_location,omitempty" gorm:"foreignKey:LocationID;references:ID"`
	MinRobins     uint8         `json:"min_robins,omitempty"`
	MaxRobins     uint8         `json:"max_robins,omitempty"`
	CreatedBy     int64         `json:"-"`
	UpdatedBy     int64         `json:"-" gorm:"-"`
	CreatedAt     time.Time     `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt     time.Time     `json:"-" gorm:"-"`
}

type CreateEventResponse struct {
	Status  *Status       `json:"status,omitempty"`
	EventId nanoid.NanoID `json:"event_id,omitempty"`
}

type GetEventResponse struct {
	Status *Status `json:"status,omitempty"`
	Event  *Event  `json:"event,omitempty"`
}

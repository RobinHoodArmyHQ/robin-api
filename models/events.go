package models

import (
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
	ID              int64     `json:"-" gorm:"primaryKey"`
	EventId         string    `json:"event_id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Description     string    `json:"description,omitempty"`
	StartTime       time.Time `json:"start_time,omitempty"`
	EventType       EventType `json:"event_type,omitempty"`
	EventLocationID int64     `json:"-"`
	EventLocation   *Location `json:"event_location,omitempty" gorm:"foreignKey:EventLocationID;references:LocationId"`
	MinRobins       uint8     `json:"min_robins,omitempty"`
	MaxRobins       uint8     `json:"max_robins,omitempty"`
	CreatedBy       int64     `json:"-"`
	UpdatedBy       int64     `json:"-"`
	CreatedAt       time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

type CreateEventResponse struct {
	Status  *Status `json:"status,omitempty"`
	EventId string  `json:"event_id,omitempty"`
}

type GetEventResponse struct {
	Status *Status `json:"status,omitempty"`
	Event  *Event  `json:"event,omitempty"`
}

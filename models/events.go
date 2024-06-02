package models

import "github.com/google/uuid"

type EventType uint8

const (
	MEAL_DRIVE EventType = 0
	ACADEMY
	OTHER
)

type VehicleType uint8

const (
	TWO_WHEELER VehicleType = 0
	FOUR_WHEELER
)

type Event struct {
	EventId              uuid.UUID `json:"event_id,omitempty"`
	Name                 string    `json:"name,omitempty"`
	Description          string    `json:"description,omitempty"`
	Timestamp            int64     `json:"timestamp,omitempty"`
	EventType            EventType `json:"event_type,omitempty"`
	PickupLocation       Location  `json:"pickup_location,omitempty"`
	DistributionLocation Location  `json:"distribution_location,omitempty"`
	AcademyLocation      Location  `json:"academy_location,omitempty"`
	MinRobins            uint8     `json:"min_robins,omitempty"`
	MaxRobins            uint8     `json:"max_robins,omitempty"`
}

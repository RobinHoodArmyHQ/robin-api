package models

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

type Event struct {
	EventId              int64     `json:"event_id,omitempty"`
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

type CreateEventResponse struct {
	Status  *Status `json:"status,omitempty"`
	EventId int64   `json:"event_id,omitempty"`
}

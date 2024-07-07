package models

import "github.com/google/uuid"

type CheckIn struct {
	ID                uint64    `json:"-" gorm:"primaryKey"`
	CheckInID         uuid.UUID `json:"check_in_id,omitempty"`
	UserID            uuid.UUID `json:"user_id,omitempty"`
	EventID           uuid.UUID `json:"event_id,omitempty"`
	PhotoURLs         []string  `json:"photo_urls,omitempty" gorm:"serializer:json"`
	Description       string    `json:"description,omitempty"`
	NoOfPeopleServed  uint64    `json:"no_of_people_served,omitempty"`
	NoOfStudentTaught uint64    `json:"no_of_student_taught,omitempty"`
	CreatedAt         string    `json:"created_at,omitempty"`
	UpdatedAt         string    `json:"updated_at,omitempty"`
}

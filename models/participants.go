package models

import (
	"time"
)

type ParticipantStatus string
type ParticipantRole string

const (
	StatusInterested ParticipantStatus = "INTERESTED"
	StatusNotGoing   ParticipantStatus = "NOT_GOING"
	StatusCheckedIn  ParticipantStatus = "CHECKED_IN"

	RoleVolunteer ParticipantRole = "VOLUNTEER"
)

type Participant struct {
	ID        uint64            `json:"-" gorm:"primaryKey"`
	UserID    uint64            `json:"-"`
	User      *User             `json:"user,omitempty" gorm:"foreignKey:ID;references:UserID"`
	EventID   int64             `json:"-"`
	Event     *Event            `json:"event,omitempty" gorm:"foreignKey:ID;references:EventID"`
	Status    ParticipantStatus `json:"status,omitempty" gorm:"type:varchar(20);default:'INTERESTED'"`
	Role      ParticipantRole   `json:"role,omitempty"`
	CreatedAt time.Time         `json:"created_at,omitempty"`
	UpdatedAt time.Time         `json:"updated_at,omitempty"`
}

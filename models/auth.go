package models

import "github.com/google/uuid"

type MobileNumber struct {
	CountryCode  uint8  `json:"country_code,omitempty"`
	MobileNumber uint64 `json:"mobile_number,omitempty"`
}

type AuthResponse struct {
	Status    *Status
	RequestID uuid.UUID `json:"request_id,omitempty"`
}

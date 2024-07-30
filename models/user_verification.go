package models

import (
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type UserVerfication struct {
	ID             uint64                 `json:"-" gorm:"primaryKey"`
	UserID         nanoid.NanoID          `json:"user_id,omitempty"`
	EmailId        string                 `json:"email_id"`
	Otp            uint64                 `json:"otp,omitempty"`
	OtpGeneratedAt time.Time              `json:"otp_generated_at,omitempty"`
	OtpExpiresAt   time.Time              `json:"otp_expires_at,omitempty"`
	OtpRetryCount  uint64                 `json:"otp_retry_count,omitempty"`
	IsVerified     int8                   `json:"is_verified,omitempty"`
	ExtraDetails   map[string]interface{} `json:"extra_details,omitempty" gorm:"serializer:json"`
}

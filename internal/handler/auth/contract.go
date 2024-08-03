package auth

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
)

type RegisterUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	EmailId   string `json:"email_id" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type RegisterUserResponse struct {
	Status    *models.Status `json:"status,omitempty"`
	IsNewUser uint8          `json:"is_new_user,omitempty"`
	UserID    string         `json:"user_id,omitempty"`
}

type LoginUserRequest struct {
	EmailId  string `json:"email_id" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Status *models.Status `json:"status,omitempty"`
	Token  string         `json:"token,omitempty"`
}

type VerifyOtpRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Otp    uint64 `json:"otp" binding:"required"`
}

type VerifyOtpResponse struct {
	Token  string         `json:"token,omitempty"`
	Status *models.Status `json:"status,omitempty"`
}

type ResendOtpRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type ResendOtpResponse struct {
	Status *models.Status `json:"status,omitempty"`
}

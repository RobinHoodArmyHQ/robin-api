package auth

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type RegisterUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	EmailId   string `json:"email_id" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type RegisterUserResponse struct {
	Status    models.Status
	IsNewUser bool          `json:"is_new_user,omitempty"`
	UserID    nanoid.NanoID `json:"user_id,omitempty"`
}

type LoginUserRequest struct {
	EmailId  string `json:"email_id" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Status models.Status
	Token  string `json:"token,omitempty"`
}

type VerifyOtpRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Otp    uint64 `json:"otp" binding:"required"`
}

type VerifyOtpResponse struct {
	Token  string `json:"token,omitempty"`
	Status models.Status
}

type ResendOtpRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type ResendOtpResponse struct {
	Status models.Status
}

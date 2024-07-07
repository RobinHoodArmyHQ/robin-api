package user

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/google/uuid"
)

type User interface {
	CreateUser(req *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
}

type CreateUserRequest struct {
	User *models.User `json:"user"`
}

type CreateUserResponse struct {
	UserID uuid.UUID `json:"user_id"`
}

type GetUserRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type GetUserResponse struct {
	User *models.User `json:"user"`
}

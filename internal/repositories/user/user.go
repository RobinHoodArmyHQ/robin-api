package user

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
)

type User interface {
	CreateUser(req *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
}

type CreateUserRequest struct {
	User *models.User `json:"user"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}

type GetUserRequest struct {
	UserID string `json:"user_id"`
}

type GetUserResponse struct {
	User *models.User `json:"user"`
}

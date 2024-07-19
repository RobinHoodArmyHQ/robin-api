package user

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type User interface {
	CreateUser(req *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
}

type CreateUserRequest struct {
	User *models.User `json:"user"`
}

type CreateUserResponse struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type GetUserRequest struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type GetUserResponse struct {
	User *models.User `json:"user"`
}

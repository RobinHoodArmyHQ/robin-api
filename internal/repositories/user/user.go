package user

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type User interface {
	CreateUser(req *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
	GetUserByEmail(req *GetUserByEmailRequest) (*GetUserByEmailResponse, error)
	UpdateUser(req *UpdateUserRequest) (*UpdateUserResponse, error)
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

type UpdateUserRequest struct {
	UserID nanoid.NanoID          `json:"user_id"`
	Values map[string]interface{} `json:"values"`
}

type UpdateUserResponse struct {
	Users *[]models.User `json:"users"`
}

type GetUserByEmailRequest struct {
	EmailId string `json:"email_id"`
}

type GetUserByEmailResponse struct {
	User *models.User `json:"user"`
}

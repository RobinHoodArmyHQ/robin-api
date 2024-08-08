package userverification

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type UserVerification interface {
	GetUserByUserID(req *GetUserByUserIdRequest) (*GetUserByUserIdResponse, error)
	CreateUser(req *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(req *UpdateUserRequest) (*UpdateUserResponse, error)
}

type GetUserByUserIdRequest struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type GetUserByUserIdResponse struct {
	User *models.UserVerification `json:"user"`
}

type CreateUserRequest struct {
	User *models.UserVerification `json:"user"`
}

type CreateUserResponse struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type UpdateUserRequest struct {
	UserID nanoid.NanoID          `json:"user_id"`
	Values map[string]interface{} `json:"values"`
}

type UpdateUserResponse struct {
	Users []*models.UserVerification
}

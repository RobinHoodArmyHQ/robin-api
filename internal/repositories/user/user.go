package user

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type User interface {
	CreateUser(req *CreateUserRequest) (*CreateUserResponse, error)
	CreateUnverifiedUser(req *CreateUnverifiedUserRequest) (*CreateUnverifiedUserResponse, error)
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
	GetUserByEmail(req *GetUserByEmailRequest) (*GetUserByEmailResponse, error)
	GetUnverifiedUserByUserID(req *GetUnverifiedUserByUserIdRequest) (*GetUnverifiedUserByUserIdResponse, error)
	UpdateUnverifiedUser(req *UpdateUnverifiedUserRequest) (*UpdateUnverifiedUserResponse, error)
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

type GetUserByEmailRequest struct {
	EmailId string `json:"email_id"`
}

type GetUserByEmailResponse struct {
	User *models.User `json:"user"`
}

type GetUnverifiedUserByUserIdRequest struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type GetUnverifiedUserByUserIdResponse struct {
	User *models.UserVerification `json:"user"`
}

type CreateUnverifiedUserRequest struct {
	User *models.UserVerification `json:"user"`
}

type CreateUnverifiedUserResponse struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type UpdateUnverifiedUserRequest struct {
	UserID nanoid.NanoID          `json:"user_id"`
	Values map[string]interface{} `json:"values"`
}

type UpdateUnverifiedUserResponse struct {
	Users *[]models.UserVerification
}

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
	GetUserByUserID(req *GetUserByUserIdRequest) (*GetUserByUserIdResponse, error)
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

type GetUserByEmailRequest struct {
	EmailId string `json:"email_id"`
}

type GetUserByEmailResponse struct {
	User *models.User `json:"user"`
}

type GetUserByUserIdRequest struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type GetUserByUserIdResponse struct {
	User *models.UserVerfication `json:"user"`
}

type CreateUnverifiedUserRequest struct {
	User *models.UserVerfication `json:"user"`
}

type CreateUnverifiedUserResponse struct {
	UserID nanoid.NanoID `json:"user_id"`
}

type UpdateUserRequest struct {
	UserID nanoid.NanoID          `jspn:"user_id"`
	Values map[string]interface{} `json:"values"`
}

type UpdateUserResponse struct {
	Users *[]models.UserVerfication
}

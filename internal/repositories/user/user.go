package user

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type User interface {
	CreateUser(req *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
	GetUserByEmailId(req *GetUserByEmailIdRequest) (*GetUserByEmailIdResponse, error)
	CheckIfUserExists(req *CheckIfUserExistsRequest) (*CheckIfUserExistsResponse, error)
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

type RegisterUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	EmailId  string `json:"email_id" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserResponse struct {
	Status    models.Status
	IsNewUser bool `json:"is_new_user,omitempty"`
}

type LoginUserRequest struct {
	EmailId  string `json:"email_id" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Status models.Status
	Token  string `json:"token,omitempty"`
}

type GetUserByEmailIdRequest struct {
	EmailId string `json:"email_id"`
}

type GetUserByEmailIdResponse struct {
	User *models.User `json:"user"`
}

type CheckIfUserExistsRequest struct {
	EmailId string `json:"email_id"`
}

type CheckIfUserExistsResponse struct {
	IsExisting bool `json:"is_existing"`
}

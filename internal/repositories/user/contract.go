package user

import "github.com/RobinHoodArmyHQ/robin-api/models"

type CreateUserRequest struct {
	User *models.User `json:"user"`
}

type CreateUserResponse struct {
	UserID uint64 `json:"user_id"`
}

type GetUserRequest struct {
	UserID uint64 `json:"user_id"`
}

type GetUserResponse struct {
	User *models.User `json:"user"`
}

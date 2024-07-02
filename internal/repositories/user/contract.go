package user

import "github.com/RobinHoodArmyHQ/robin-api/models"

type CreateUserRequest struct {
	User *models.User
}

type CreateUserResponse struct {
	UserID int64
}

type GetUserRequest struct {
	UserID int64
}

type GetUserResponse struct {
	User *models.User
}

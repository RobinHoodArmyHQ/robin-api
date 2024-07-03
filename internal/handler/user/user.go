package user

import (
	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/contract"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "user_id is required"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "invalid user id"})
		return
	}

	userRepo := env.FromContext(c).UserRepository
	userResponse, err := userRepo.GetUser(&user.GetUserRequest{UserID: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &contract.Response{Message: err.Error()})
		return
	}

	if userResponse == nil || userResponse.User == nil {
		c.JSON(http.StatusNotFound, &contract.Response{Message: "user not found"})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

func CreateUserHandler(c *gin.Context) {
	userReq := &user.CreateUserRequest{}
	err := c.BindJSON(userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: err.Error()})
		return
	}

	if userReq.User == nil {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "user is required"})
		return
	}

	userReq.User.UserId = 0
	userRepo := env.FromContext(c).UserRepository
	userResponse, err := userRepo.CreateUser(userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &contract.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

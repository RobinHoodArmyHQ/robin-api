package checkin

import (
	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/contract"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/checkin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCheckInHandler(c *gin.Context) {
	checkIn := &checkin.CreateCheckInRequest{}
	err := c.BindJSON(checkIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: err.Error()})
		return
	}

	if checkIn.CheckIn == nil {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "check_in is required"})
		return
	}

	checkIn.CheckIn.CheckInID = 0
	checkInRepo := env.FromContext(c).CheckInRepository
	checkInResp, err := checkInRepo.CreateCheckIn(checkIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &contract.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, checkInResp)
}

func GetCheckInHandler(c *gin.Context) {
	checkInIDStr := c.Param("check_in_id")
	if checkInIDStr == "" {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "check_in_id is required"})
		return
	}

	checkInID, err := strconv.ParseUint(checkInIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "invalid check_in id"})
		return
	}

	checkInRepo := env.FromContext(c).CheckInRepository
	checkInResp, err := checkInRepo.GetCheckIn(&checkin.GetCheckInRequest{CheckInID: checkInID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &contract.Response{Message: err.Error()})
		return
	}

	if checkInResp == nil || checkInResp.CheckIn == nil {
		c.JSON(http.StatusNotFound, &contract.Response{Message: "check_in not found"})
		return
	}

	c.JSON(http.StatusOK, checkInResp)
}

func GetUserCheckInsHandler(c *gin.Context) {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "user_id is required"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "invalid user_id"})
		return
	}

	checkInRepo := env.FromContext(c).CheckInRepository
	checkInResp, err := checkInRepo.GetUserCheckIns(&checkin.GetUserCheckInsRequest{UserID: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &contract.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, checkInResp)
}

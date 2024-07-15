package checkin

import (
	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/contract"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/checkin"
	"github.com/gin-gonic/gin"
	"net/http"
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

	checkInRepo := env.FromContext(c).CheckInRepository
	checkInResp, err := checkInRepo.CreateCheckIn(checkIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &contract.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, checkInResp)
}

func GetCheckInHandler(c *gin.Context) {
	checkInID := c.Param("check_in_id")
	if checkInID == "" {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "check_in_id is required"})
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
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, &contract.Response{Message: "user_id is required"})
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

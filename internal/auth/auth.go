package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthHandler(c *gin.Context) {

	// validate country isd code
	countryCode, err := strconv.ParseUint(c.PostForm("country_code"), 10, 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AuthResponse{
			Status: models.StatusFailed(fmt.Sprintf("invalid country code %d", countryCode)),
		})
		return
	}

	// validate mobile number
	mobileNumber, err := strconv.ParseUint(c.PostForm("mobile_number"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AuthResponse{
			Status: models.StatusFailed(fmt.Sprintf("invalid mobile number %d", mobileNumber)),
		})
		return
	}

	// generate request id and send response
	requestId := uuid.Must(uuid.NewRandom())
	c.JSON(http.StatusOK, models.AuthResponse{
		Status:    models.StatusSuccess(),
		RequestID: requestId,
	})
}

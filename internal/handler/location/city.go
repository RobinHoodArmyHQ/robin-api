package location

import (
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/gin-gonic/gin"
)

func GetCitiesHandler(c *gin.Context) {
	ev := env.FromContext(c)

	resp, err := ev.LocationRepository.GetCities(&repositories.GetCitiesRequest{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &models.GetCitiesResponse{
		Status: models.StatusSuccess(),
		Cities: resp.Cities,
	})
}

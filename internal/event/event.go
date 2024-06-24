package event

import (
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/gin-gonic/gin"
)

func CreateEventHandler(c *gin.Context) {
	ev := env.FromContext(c)

	var req *models.Event
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed("Invalid request"))
		return
	}

	if err = validateCreateEventRequest(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed(err.Error()))
		return
	}

	resp, err := ev.EventRepository.CreateEvent(&event.CreateEventRequest{Event: req})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &models.CreateEventResponse{
		Status:  models.StatusSuccess("Event created successfully"),
		EventId: resp.EventID,
	})
}

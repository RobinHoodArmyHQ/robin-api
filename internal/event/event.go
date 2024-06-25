package event

import (
	"net/http"
	"strconv"

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

func GetEventHandler(c *gin.Context) {
	ev := env.FromContext(c)

	eventIDStr := c.Param("event_id")
	if eventIDStr == "" || eventIDStr == "0" {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed("Missing event id"))
		return
	}

	eventID, err := strconv.ParseInt(eventIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed("Invalid event id"))
		return
	}

	resp, err := ev.EventRepository.GetEvent(&event.GetEventRequest{EventID: eventID})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &models.GetEventResponse{
		Status: models.StatusSuccess("Event fetched successfully"),
		Event:  resp.Event,
	})
}

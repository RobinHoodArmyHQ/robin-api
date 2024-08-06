package event

import (
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
	eventsvc "github.com/RobinHoodArmyHQ/robin-api/internal/services/event"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/gin-gonic/gin"
)

func CreateEventHandler(c *gin.Context) {
	ev := env.FromContext(c)

	var req *models.Event
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed(err.Error()))
		return
	}

	if err = validateCreateEventRequest(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed(err.Error()))
		return
	}

	cityResp, err := ev.LocationRepository.GetCityByID(&repositories.GetCityByCityIDRequest{CityID: req.EventLocation.City.CityId})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	req.EventLocation.CityId = cityResp.City.ID
	req.EventLocation.City = cityResp.City
	resp, err := ev.EventRepository.CreateEvent(&event.CreateEventRequest{Event: req})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &models.CreateEventResponse{
		Status:  models.StatusSuccess(),
		EventId: resp.EventID,
	})
}

func GetEventHandler(c *gin.Context) {
	ev := env.FromContext(c)

	eventID := c.Param("event_id")
	if eventID == "" || eventID == "0" {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed("Missing event id"))
		return
	}

	resp, err := ev.EventRepository.GetEvent(&event.GetEventRequest{EventID: nanoid.NanoID(eventID)})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &models.GetEventResponse{
		Status: models.StatusSuccess(),
		Event:  resp.Event,
	})
}

// GetEventsHandler return list of events based on user location and filter
func GetEventsHandler(c *gin.Context) {
	ev := env.FromContext(c)
	req := &GetEventsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed(err.Error()))
		return
	}

	cityResp, err := ev.LocationRepository.GetCityByID(&repositories.GetCityByCityIDRequest{CityID: req.CityId})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	eventSvc := eventsvc.New()
	resp, err := eventSvc.GetEventFeed(c, &eventsvc.GetEventFeedRequest{
		Page:   req.Page,
		Limit:  req.Limit,
		CityId: cityResp.City.ID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &GetEventsResponse{
		Page:   req.Page,
		Limit:  req.Limit,
		Events: resp.Events,
	})
}

func InterestedEventHandler(c *gin.Context) {
	req := &InterestedEventRequest{}
	res := &InterestedEventResponse{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Status = models.StatusFailed(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if req.EventID == "" {
		res.Status = models.StatusFailed("Missing event id")
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	eventSvc := eventsvc.New()
	err = eventSvc.MarkEventInterested(c, req.EventID)
	if err != nil {
		res.Status = models.StatusFailed(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Status = models.StatusSuccess()
	c.JSON(http.StatusOK, res)
}

func GetParticipantsHandler(c *gin.Context) {
	res := &GetParticipantsResponse{}
	eventID := c.Param("event_id")
	if eventID == "" || eventID == "0" {
		res.Status = models.StatusFailed("Missing event id")
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	eventSvc := eventsvc.New()
	resp, err := eventSvc.GetEventParticipants(c, nanoid.NanoID(eventID))
	if err != nil {
		res.Status = models.StatusFailed(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Status = models.StatusSuccess()
	res.Participants = resp.Participants
	c.JSON(http.StatusOK, res)
}

package router

import (
	"context"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/checkin"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/user"
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/auth"
	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/event"
	"github.com/RobinHoodArmyHQ/robin-api/internal/location"
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
)

func Initialize(ctx context.Context, ev *env.Env) *gin.Engine {
	r := gin.New()
	r.ContextWithFallback = true

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// compress responses using gzip
	r.Use(gzip.DefaultHandler().Gin)

	// add env to the context
	r.Use(env.MiddleWare(ev))

	// health check route
	r.GET("/health", HealthcheckHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/cities", location.GetCitiesHandler)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("", auth.AuthHandler)
	}

	eventGroup := r.Group("/event")
	eventGroup.Use(isUserLoggedIn)
	setupEventGroup(eventGroup)

	userGroup := r.Group("/user")
	userGroup.Use(isUserLoggedIn)
	{
		userGroup.GET("/:user_id", user.GetUserHandler)
		// TODO: this is for testing purposes only - create user at auth level
		userGroup.POST("/", user.CreateUserHandler)
	}

	checkInGroup := r.Group("/checkin")
	checkInGroup.Use(isUserLoggedIn)
	{
		checkInGroup.POST("/", checkin.CreateCheckInHandler)
		checkInGroup.GET("/:check_in_id", checkin.GetCheckInHandler)
		// Check-in list for a user
		checkInGroup.GET("/list", checkin.GetUserCheckInsHandler)
	}

	return r
}

func setupEventGroup(eventGroup *gin.RouterGroup) {
	eventGroup.GET("/:event_id", event.GetEventHandler)
	eventGroup.POST("/", isAdminUser, event.CreateEventHandler)
}

func HealthcheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "healthy")
}

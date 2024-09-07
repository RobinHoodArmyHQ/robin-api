package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/auth"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/checkin"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/event"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/location"
	"github.com/RobinHoodArmyHQ/robin-api/internal/handler/user"
	"github.com/RobinHoodArmyHQ/robin-api/internal/services/photo"
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

	{
		r.POST("/auth", auth.AuthHandler)
		r.POST("/auth/signup", auth.RegisterUser)
		r.POST("/auth/verify", auth.VerifyOtp)
		r.POST("/auth/verify/resend", auth.ResendOtp)
		r.POST("/auth/login", auth.LoginUser)
		r.POST("/auth/sendPasswordResetLink", auth.SendPasswordResetLink)
		r.POST("/auth/resetPassword", auth.ResetPassword)
	}

	r.Use(isUserLoggedIn)
	{
		r.GET("/events", event.GetEventsHandler)
		r.POST("/photo", photo.PhotoUploadHandler)
	}

	eventGroup := r.Group("/event")
	eventGroup.Use(isUserLoggedIn)
	{
		eventGroup.GET("/:event_id", event.GetEventHandler)
		eventGroup.POST("/", isAdminUser, event.CreateEventHandler)
		eventGroup.POST("/interested", event.InterestedEventHandler)
		eventGroup.GET("/:event_id/participants", event.GetParticipantsHandler)
		eventGroup.POST("/:event_id/checkin", checkin.CreateCheckInHandler)
	}

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
		checkInGroup.GET("/:checkin_id", checkin.GetCheckInHandler)
		// Check-in list for a user
		checkInGroup.GET("/list", checkin.GetUserCheckInsHandler)
	}

	return r
}

func HealthcheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "healthy")
}

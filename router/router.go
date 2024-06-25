package router

import (
	"context"
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/auth"
	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/event"
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

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("", auth.AuthHandler)
	}

	eventGroup := r.Group("/event")
	eventGroup.Use(isUserLoggedIn)
	setupEventGroup(eventGroup)

	return r
}

func setupEventGroup(eventGroup *gin.RouterGroup) {
	eventGroup.GET("/:event_id", event.GetEventHandler)
	eventGroup.POST("/", isAdminUser, event.CreateEventHandler)
}

func HealthcheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "healthy")
}

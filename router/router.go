package router

import (
	"context"
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
)

func Initialize(ctx context.Context) *gin.Engine {
	r := gin.New()
	r.ContextWithFallback = true

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// compress responses using gzip
	r.Use(gzip.DefaultHandler().Gin)

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

	return r
}

func HealthcheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "healthy")
}

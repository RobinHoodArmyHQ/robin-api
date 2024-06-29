package router

import (
	"github.com/gin-gonic/gin"
)

func isUserLoggedIn(c *gin.Context) {
	c.Next()
}

func isAdminUser(c *gin.Context) {
	c.Next()
}

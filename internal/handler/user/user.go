package user

import "github.com/gin-gonic/gin"

func GetUserHandler(c *gin.Context) {
	c.String(200, "User handler")
}

func CreateUserHandler(c *gin.Context) {
	c.String(200, "Create user handler")
}

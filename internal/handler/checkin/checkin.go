package checkin

import "github.com/gin-gonic/gin"

func CreateCheckInHandler(c *gin.Context) {
	c.String(200, "Create checkin handler")
}

func GetCheckInHandler(c *gin.Context) {
	c.String(200, "Get checkin handler")
}

func GetUserCheckInsHandler(c *gin.Context) {
	c.String(200, "Get user checkins handler")
}

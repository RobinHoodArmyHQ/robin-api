package router

import (
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/util"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
	"github.com/gin-gonic/gin"
)

func isUserLoggedIn(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "request does not contain an access token",
		})
		c.Abort()
		return
	}

	claims, err := util.VerifyJwt(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("user_id", nanoid.NanoID(claims.UserInfo["user_id"].(string)))
	c.Set("user_role", claims.UserInfo["user_role"].(string))
	c.Next()
}

func isAdminUser(c *gin.Context) {
	c.Next()
}

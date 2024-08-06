package router

import (
	"github.com/RobinHoodArmyHQ/robin-api/pkg/ctxmeta"
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/util"
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

	userId := claims.UserInfo["user_id"].(string)
	role := claims.UserInfo["user_role"].(string)
	ctxmeta.SetUser(c, userId)
	ctxmeta.SetRole(c, role)
	c.Next()
}

func isAdminUser(c *gin.Context) {
	c.Next()
}

package ctxmeta

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

const (
	UserKey = "userId"
	RoleKey = "role"
)

func set(c *gin.Context, key string, value interface{}) {
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, key, value)
	c.Request = c.Request.WithContext(ctx)
	c.Set(key, value)
}

func getStringValue(ctx context.Context, key string) string {
	value := ctx.Value(key)
	if value == nil {
		return ""
	}

	return value.(string)
}

func SetUser(c *gin.Context, userId string) {
	set(c, UserKey, userId)
}

func SetRole(c *gin.Context, role string) {
	set(c, RoleKey, role)
}

func GetUser(ctx context.Context) nanoid.NanoID {
	return nanoid.NanoID(getStringValue(ctx, UserKey))
}

func GetRole(ctx context.Context) string {
	return getStringValue(ctx, RoleKey)
}

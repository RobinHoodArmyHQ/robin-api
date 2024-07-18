package env

import (
	"context"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/checkin"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/user"

	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/gin-gonic/gin"
)

const (
	EnvCtxKey = "env"
)

type Env struct {
	SqlDBConn *database.SqlDB

	EventRepository    event.EventRepository
	UserRepository     user.User
	CheckInRepository  checkin.CheckIn
	LocationRepository repositories.LocationRepository
}

func FromContext(ctx context.Context) *Env {
	env, ok := ctx.Value(EnvCtxKey).(*Env)
	if !ok {
		panic("could not fetch env from context")
	}

	return env
}

func MiddleWare(ev *Env) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Set(EnvCtxKey, ev)
		c.Next()
	}
}

func NewEnv(options ...func(env *Env)) *Env {
	env := &Env{}

	for _, option := range options {
		option(env)
	}

	return env
}

func WithSqlDBConn(db *database.SqlDB) func(*Env) {
	return func(env *Env) {
		env.SqlDBConn = db
	}
}

func WithEventRepository(eventRepo event.EventRepository) func(*Env) {
	return func(env *Env) {
		env.EventRepository = eventRepo
	}
}

func WithUserRepository(userRepo user.User) func(*Env) {
	return func(env *Env) {
		env.UserRepository = userRepo
	}
}

func WithCheckInRepository(checkInRepo checkin.CheckIn) func(*Env) {
	return func(env *Env) {
		env.CheckInRepository = checkInRepo
	}
}

func WithLocationRepository(repo repositories.LocationRepository) func(*Env) {
	return func(env *Env) {
		env.LocationRepository = repo
	}
}

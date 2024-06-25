package env

import (
	"context"
	"database/sql"

	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
	"github.com/gin-gonic/gin"
)

const (
	EnvCtxKey = "env"
)

type Env struct {
	SqlDBConn *sql.DB

	EventRepository event.EventRepository
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

func WithSqlDBConn(db *sql.DB) func(*Env) {
	return func(env *Env) {
		env.SqlDBConn = db
	}
}

func WithEventRepository(eventRepo event.EventRepository) func(*Env) {
	return func(env *Env) {
		env.EventRepository = eventRepo
	}
}

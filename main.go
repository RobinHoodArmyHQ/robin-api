package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	sqlRepos "github.com/RobinHoodArmyHQ/robin-api/internal/repositories/sql"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/RobinHoodArmyHQ/robin-api/router"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Initialize DB
	dbConn, err := database.Connect(logger)
	if err != nil {
		logger.Fatal("could not connect to database, err: %v", zap.Error(err))
	}
	logger.Info("connected to database")
	defer dbConn.Close()

	ev := env.NewEnv(
		env.WithSqlDBConn(dbConn),
		env.WithEventRepository(sqlRepos.NewEventRepository(logger, dbConn)),
	)

	r := router.Initialize(ctx, ev)
	srv := &http.Server{
		Addr:         ":" + viper.GetString("listen.port"),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

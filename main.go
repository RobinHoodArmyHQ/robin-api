package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/router"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/net/context"
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

	r := router.Initialize(ctx)
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

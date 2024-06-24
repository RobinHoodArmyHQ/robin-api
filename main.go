package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	sqlRespo "github.com/RobinHoodArmyHQ/robin-api/internal/repositories/sql"
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

	// Format: username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", viper.GetString("mysql.username"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetString("mysql.port"), viper.GetString("mysql.database"))

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Ping the database to ensure a successful connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	logger.Info("Successfully connected to the database!")

	// Close the database connection when main function exits
	defer db.Close()

	ev := env.NewEnv(
		env.WithSqlDB(db),
		env.WithEventRepository(sqlRespo.NewEventRepository(db)),
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

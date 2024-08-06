package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/sql"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/sql/checkin"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/sql/event"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/sql/participants"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/sql/user"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"github.com/RobinHoodArmyHQ/robin-api/router"
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
		env.WithEventRepository(event.NewEventRepository(logger, dbConn)),
		env.WithParticipantsRepository(participants.NewParticipantsRepository(logger, dbConn)),
		env.WithUserRepository(user.New(logger, dbConn)),
		env.WithCheckInRepository(checkin.New(logger, dbConn)),
		env.WithLocationRepository(sql.NewLocationRepository(logger, dbConn)),
		env.WithPhotoRepository(sql.NewPhotoRepository(logger, dbConn)),
		env.WithS3Service(initializeS3()),
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

func initializeS3() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(viper.GetString("s3.access_key_id"), viper.GetString("s3.secret_access_key"), ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(viper.GetString("s3.region")),
		Endpoint:         aws.String(viper.GetString("s3.endpoint")),
	})

	if err != nil {
		panic(fmt.Errorf("error creating aws session: %w", err))
	}

	return s3.New(sess)
}

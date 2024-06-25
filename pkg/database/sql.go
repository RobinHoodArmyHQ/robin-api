package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func getDSN() string {
	// Format: username:password@protocol(address)/dbname?param=value
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true&charset=utf8&parseTime=true&timeout=5s",
		viper.GetString("mysql.username"), viper.GetString("mysql.password"), viper.GetString("mysql.host"),
		viper.GetString("mysql.port"), viper.GetString("mysql.database"),
	)
}

// Connect will open a connection to the database
func Connect(logger *zap.Logger) (*sql.DB, error) {
	db, err := sql.Open("mysql", getDSN())
	if err != nil {
		return nil, err
	}

	var retries = 30

	// Ping the database to ensure a successful connection
	for retries > 0 {
		err = db.Ping()

		retries--

		if err != nil {
			time.Sleep(time.Second * 1)
		} else {
			break
		}

		if retries == 0 {
			return nil, errors.New("could not connect to database after max retries")
		}
	}

	return db, nil
}

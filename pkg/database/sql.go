package database

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type SqlDB struct {
	logger *zap.Logger
	master *connection
}

func (db *SqlDB) Master() *connection {
	return db.master
}

func (db *SqlDB) Close() {
	gdb, err := db.master.DB.DB()
	if err != nil {
		db.logger.Error("could not close database connection", zap.Error(err))
		return
	}
	gdb.Close()
}

// Connect will open a connection to the database
func Connect(logger *zap.Logger) (*SqlDB, error) {
	db := &SqlDB{
		logger: logger,
	}

	// Initialize master DB
	masterDBConn, err := NewConnection(
		User(viper.GetString("mysql.username")),
		Password(viper.GetString("mysql.password")),
		Host(viper.GetString("mysql.host")),
		Port(viper.GetInt("mysql.port")),
		Schema(viper.GetString("mysql.database")),
		Debug(viper.GetBool("mysql.debug")),
		MaxOpenConnections(viper.GetInt("mysql.max_open_connections")),
		MaxIdleConnections(viper.GetInt("mysql.max_idle_connections")),
		ConnectionMaxLifetime(viper.GetDuration("mysql.conn_max_lifetime")),
	)
	if err != nil {
		return nil, err
	}

	db.master = masterDBConn

	if err = initTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initTables(db *SqlDB) error {
	// Initialize tables
	err := db.master.DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	err = db.master.DB.AutoMigrate(&models.Event{})
	if err != nil {
		return err
	}

	err = db.master.DB.AutoMigrate(&models.CheckIn{})
	if err != nil {
		return err
	}

	return nil
}

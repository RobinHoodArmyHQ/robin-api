package database

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type connection struct {
	user     string
	pass     string
	host     string
	port     int
	database string
	dbParams string

	debug              bool
	maxOpenConnections int
	maxIdleConnections int
	connMaxLifetime    time.Duration

	*gorm.DB
}

func getDSN(conn *connection) string {
	// Format: username:password@protocol(address)/dbname?param=value
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		conn.user,
		conn.pass,
		conn.host,
		conn.port,
		conn.database,
		conn.dbParams,
	)
}

// NewConnection returns a new database connection, use option functions to
// configure connection properties
func NewConnection(options ...func(*connection) error) (*connection, error) {
	conn := &connection{
		host:               "127.0.0.1",
		port:               3306,
		maxOpenConnections: 100,
		maxIdleConnections: 10,
		connMaxLifetime:    time.Hour,
		debug:              true,
		dbParams:           "allowNativePasswords=true&charset=utf8&parseTime=true&timeout=5s",
	}

	for _, option := range options {
		err := option(conn)

		if err != nil {
			return nil, err
		}
	}

	dsn := getDSN(conn)

	gconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := gconn.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(conn.maxIdleConnections)
	db.SetMaxOpenConns(conn.maxOpenConnections)
	db.SetConnMaxLifetime(conn.connMaxLifetime)

	if conn.debug {
		gconn = gconn.Debug()
	}

	var retries = 30

	// Open doesn't actually open a connection. To be sure we can connect,
	// try to ping the server. This is usually required during development
	// when the DB container takes a while to initialise
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

	conn.DB = gconn

	return conn, nil
}

// User sets the username to be used for the connection while calling NewConnection
func User(user string) func(*connection) error {
	return func(c *connection) error {
		c.user = user
		return nil
	}
}

// Password sets the password to be used for the connection while calling NewConnection
func Password(passwd string) func(*connection) error {
	return func(c *connection) error {
		c.pass = passwd
		return nil
	}
}

// Host sets the hostname to be used while calling NewConnection, default localhost
func Host(host string) func(*connection) error {
	return func(c *connection) error {
		c.host = host
		return nil
	}
}

// Port sets the port to connect while calling NewConnection, defaults to 3306
func Port(port int) func(*connection) error {
	return func(c *connection) error {
		c.port = port
		return nil
	}
}

// Schema sets the database name while calling NewConnection
func Schema(db string) func(*connection) error {
	return func(c *connection) error {
		c.database = db
		return nil
	}
}

// Debug enables or disables the debug level for database
// With debug=on, all queries are logged
func Debug(enable bool) func(*connection) error {
	return func(c *connection) error {
		c.debug = enable
		return nil
	}
}

// MaxOpenConnections sets the max number of connections that can be made to the database
// default 100
//
// https://golang.org/pkg/database/sql/#DB.SetMaxOpenConns
func MaxOpenConnections(n int) func(*connection) error {
	return func(c *connection) error {
		c.maxOpenConnections = n
		return nil
	}
}

// MaxIdleConnections sets the maximum number of connections that can remain idle
// default 20
//
// https://golang.org/pkg/database/sql/#DB.SetMaxIdleConns
func MaxIdleConnections(n int) func(*connection) error {
	return func(c *connection) error {
		c.maxIdleConnections = n
		return nil
	}
}

// ConnectionMaxLifetime sets the maximum amount of time a connection may be reused.
// Expired connections may be closed lazily before reuse.
// If d <= 0, connections are reused forever.
//
// https://golang.org/pkg/database/sql/#DB.SetConnMaxLifetime
func ConnectionMaxLifetime(d time.Duration) func(*connection) error {
	return func(c *connection) error {
		c.connMaxLifetime = d
		return nil
	}
}

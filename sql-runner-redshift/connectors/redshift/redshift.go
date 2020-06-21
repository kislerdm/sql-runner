package redshift

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Config configuration to connect to database
type Config struct {
	Host       string
	Port       int
	DbName     string
	User       string
	DbPassword string
}

// Client db client
type Client struct {
	*sql.DB
}

// New function to connect to database
func New(dbConf *Config) (*Client, error) {
	connectionConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.DbPassword, dbConf.DbName)

	dbConnection, err := sql.Open("postgres", connectionConfig)
	if err != nil {
		return nil, err
	}

	err = dbConnection.Ping()
	if err != nil {
		dbConnection.Close()
		return nil, err
	}

	return &Client{dbConnection}, nil
}

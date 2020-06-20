package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config connection configuration
type Config struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsRegion          string
	Host               string
	Port               int
	dbName             string
	User               string
	dbPassword         string
}

// readEnv function to read env variables
func readEnv() (Config, error) {
	AwsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	AwsRegion := os.Getenv("AWS_DEFAULT_REGION")

	if !(AwsAccessKeyID != "" && AwsSecretAccessKey != "") {
		return Config{}, fmt.Errorf("export 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY' as env vars")
	}

	if AwsRegion == "" {
		AwsRegion = "us-west-2"
	}

	Host := os.Getenv("DB_HOST")
	if Host == "" {
		Host = "localhost"
	}

	Port := os.Getenv("DB_PORT")
	var PortInt int
	if Port == "" {
		PortInt = 5432
	} else {
		p, err := strconv.Atoi(Port)
		if err != nil {
			return Config{}, fmt.Errorf("specify int port as 'DB_PORT' env variable")
		}
		PortInt = p
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}

	User := os.Getenv("DB_USER")
	if User == "" {
		User = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}

	return Config{
		AwsAccessKeyID:     AwsAccessKeyID,
		AwsSecretAccessKey: AwsSecretAccessKey,
		AwsRegion:          AwsRegion,
		Host:               Host,
		Port:               PortInt,
		dbName:             dbName,
		User:               User,
		dbPassword:         dbPassword,
	}, nil
}

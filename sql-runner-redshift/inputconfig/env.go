package inputconfig

import (
	"fmt"
	"os"
	"strconv"
)

// EnvVars env vars configuration
type EnvVars struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsRegion          string
	Host               string
	Port               int
	dbName             string
	User               string
	dbPassword         string
}

// ReadEnvVars function to read env variables
func ReadEnvVars() (*EnvVars, error) {
	AwsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	AwsRegion := os.Getenv("AWS_DEFAULT_REGION")

	if !(AwsAccessKeyID != "" && AwsSecretAccessKey != "") {
		return &EnvVars{},
			fmt.Errorf("export 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY' as env vars")
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
			return &EnvVars{}, fmt.Errorf("specify int port as 'DB_PORT' env variable")
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

	return &EnvVars{
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

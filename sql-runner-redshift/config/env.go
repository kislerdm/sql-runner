package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kislerdm/sql-runner/sql-runner-redshift/connectors/redshift"
	"github.com/kislerdm/sql-runner/sql-runner-redshift/connectors/s3"
)

// EnvVars env vars configuration
type EnvVars struct {
	AWS *s3.Config
	DB  *redshift.Config
}

// ReadEnvVars function to read env variables
func ReadEnvVars() (*EnvVars, error) {
	AwsRegion := os.Getenv("AWS_DEFAULT_REGION")

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

	DbName := os.Getenv("DB_NAME")
	if DbName == "" {
		DbName = "postgres"
	}

	User := os.Getenv("DB_USER")
	if User == "" {
		User = "postgres"
	}

	DbPassword := os.Getenv("DB_PASSWORD")
	if DbPassword == "" {
		DbPassword = "postgres"
	}

	return &EnvVars{
		AWS: &s3.Config{
			Region: AwsRegion,
		},
		DB: &redshift.Config{
			Host:       Host,
			Port:       PortInt,
			DbName:     DbName,
			User:       User,
			DbPassword: DbPassword,
		},
	}, nil
}

package main

import (
	"log"
	"os"

	"github.com/kislerdm/sql-runner/sql-runner-redshift/config"
	db "github.com/kislerdm/sql-runner/sql-runner-redshift/connectors/redshift"
	"github.com/kislerdm/sql-runner/sql-runner-redshift/connectors/s3"
	"github.com/kislerdm/sql-runner/sql-runner-redshift/helper"
	"github.com/kislerdm/sql-runner/sql-runner-redshift/sqlparser"
)

func main() {
	// read env vars
	conf, err := config.ReadEnvVars()
	if err != nil {
		log.Fatal(err)
	}
	// parse stdin args
	args, err := config.ReadArgs()
	if err != nil {
		log.Fatal(err)
	}

	// read SQL statements
	var sqlText string
	if args.SQL != "" {
		sqlText = args.SQL
	} else if args.BucketSQL != "" && args.PathSQL != "" {
		// instantiate s3 client
		s3Client := s3.New(conf.AWS)

		// check if provided buckets are present in account
		bucketsList, err := s3Client.ListBucketNames()
		if err != nil {
			log.Fatalf("Cannot list buckets: %s.", err)
		}

		flag := false
		for _, bucket := range []string{args.BucketSQL} {
			if !helper.InArrayStr(bucketsList, bucket) {
				log.Printf("Specified bucket `%s` not found.", bucket)
				flag = true
			}
		}
		if flag {
			os.Exit(1)
		}

		// read sql queries file from s3
		sqlText, err = s3Client.ReadObjectText(args.BucketSQL, args.PathSQL)
		if err != nil {
			log.Fatalf("Cannot read sql file from 's3://%s/%s': %s",
				args.BucketSQL, args.PathSQL, err)
		}
	} else {
		log.Print("Queries aren't provided. Nothing to execute.")
	}

	// parse and impute sql paramters if provided
	if args.SQLParameters != "" {
		SQLParameters, err := helper.SQLParametersParser(args.SQLParameters)
		if err != nil {
			log.Fatalf("Cannot parse provided SQL query format paramters: %s", err)
		}
		sqlText = sqlparser.FormatQuery(sqlText, SQLParameters)
	}

	// remove comments from sql statements
	sqlText = sqlparser.StripComments(sqlText)

	// split statements to list of queries
	sqlQueries := sqlparser.SplitQueries(sqlText)

	// connect to database
	dbClient, err := db.New(conf.DB)
	if err != nil {
		log.Fatalf("Cannot connect to database: %s", err)
	}

	// execute queries
	transaction, err := dbClient.Begin()
	if err != nil {
		log.Fatalf("Cannot open transaction: %s", err)
	}

	for _, query := range sqlQueries {
		_, err := transaction.Exec(query)
		if err != nil {
			transaction.Rollback()
			dbClient.Close()
			log.Fatalf("Cannot execute query:\n%s\nError: %s", query, err)
		}
	}

	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		dbClient.Close()
		log.Fatalf("Transaction commit error: %s", err)
	}
}

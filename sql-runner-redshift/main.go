package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kislerdm/sql-runner/sql-runner-redshift/inputconfig"

	_ "github.com/lib/pq"
)

// inArray function to check if an element is present in array
func inArray(array []string, testElement string) bool {
	for _, s := range array {
		if testElement == s {
			return true
		}
	}
	return false
}

// ListBucketNames function to list buckets names in AWS account
func ListBucketNames(client *s3.S3, input *s3.ListBucketsInput) ([]string, error) {
	output := []string{}

	result, err := client.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return output, fmt.Errorf(aerr.Error())
			}
		} else {
			return output, err
		}
	}

	for _, bucket := range result.Buckets {
		output = append(output, *bucket.Name)
	}
	return output, nil
}

// readObjectText function to read sql file from s3 bucket as text
func readObjectText(client *s3.S3, bucket string, path string) (string, error) {
	s3ObjectResp, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(s3ObjectResp.Body)
	return buf.String(), nil
}

// dbConnect function to establish connection to database
func dbConnect(connectionConfig string) (*sql.DB, error) {
	dbConnection, err := sql.Open("postgres", connectionConfig)
	if err != nil {
		return nil, err
	}

	err = dbConnection.Ping()
	if err != nil {
		return nil, err
	}
	return dbConnection, nil
}

func main() {
	// read env vars
	conf, err := inputconfig.ReadEnvVars()
	if err != nil {
		log.Fatal(err)
	}
	// parse stdin args
	args, err := inputconfig.ReadArgs()
	if err != nil {
		log.Fatal(err)
	}

	// instantiate s3 client
	awsCnf := aws.Config{
		Region: aws.String(conf.AwsRegion),
	}
	sess := session.Must(session.NewSession(&awsCnf))
	svc := s3.New(sess, &awsCnf)

	// // check if provided buckets are present in account
	bucketsList, err := ListBucketNames(svc, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Cannot list buckets: %s.", err)
	}

	flag := false
	for _, bucket := range []string{args.BucketSQL, args.BucketData} {
		if !inArray(bucketsList, bucket) {
			log.Panicf("Specified bucket `%s` not found.", bucket)
			flag = true
		}
		if flag {
			os.Exit(1)
		}
	}

	// // test connection to db
	// dbConnConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	conf.Host, conf.Port, conf.User, conf.dbPassword, conf.dbName)

	// dbConnection, err := dbConnect(dbConnConfig)
	// if err != nil {
	// 	log.Fatalf("Cannot connect to database: %s", err)
	// }

	// err = dbConnection.Close()
	// if err != nil {
	// 	log.Print("Cannot close connection to database")
	// }

	// // read sql queries file from s3
	// sqlQueries, err := readObjectText(svc, args.bucketSQL, args.pathSQL)
	// if err != nil {
	// 	log.Print("Cannot read sql file from 's3://%s/%s': %s",
	// 		args.bucketSql, args.pathSql, err)
	// }
	// fmt.Print(sqlQueries)

	// var result map[string]interface{}

	// results, err := dbConnection.Query("select current_date as d;")
	// if err != nil {
	// 	log.Fatalf("SQL execution error: %s", err)
	// }
	// defer results.Close()

	// for results.Next() {
	// 	var col string
	// 	if err := results.Scan(&col); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%s", col)
	// }
	// if err := results.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}

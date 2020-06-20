package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

func main() {
	env, err := readEnv()
	if err != nil {
		log.Fatal(err)
	}

	args := getArgs()

	cnf := aws.Config{
		Region: aws.String(env.AwsRegion),
	}

	sess := session.Must(session.NewSession(&cnf))
	svc := s3.New(sess, &cnf)
	input := &s3.ListBucketsInput{}

	bucketsList, err := ListBucketNames(svc, input)
	if err != nil {
		log.Fatalf("Cannot list buckets: %s.", err)
	}

	if !inArray(bucketsList, args.bucketSql) {
		log.Fatalf("Specified bucket `%s` not found.", args.bucketSql)
	}
}

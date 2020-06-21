package s3

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Config configuration to connect to AWS infra
type Config struct {
	Region string
}

// Client s3 client
type Client struct {
	*s3.S3
}

// New function to instantiate s3 client
func New(awsCnf *Config) *Client {
	conf := aws.Config{
		Region: aws.String(awsCnf.Region),
	}
	session := session.Must(session.NewSession(&conf))
	return &Client{s3.New(session, &conf)}
}

// ListBucketNames function to list buckets names in AWS account
func (c *Client) ListBucketNames() ([]string, error) {
	output := []string{}

	result, err := c.ListBuckets(&s3.ListBucketsInput{})
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

// ReadObjectText function to read object from s3 bucket as text
func (c *Client) ReadObjectText(bucket string, path string) (string, error) {
	s3ObjectResp, err := c.GetObject(&s3.GetObjectInput{
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

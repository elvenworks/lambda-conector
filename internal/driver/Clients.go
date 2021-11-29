package driver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	cloudwatchlogsV1 "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func GetAWSCloudWatchClient(region string) (*cloudwatch.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}
	scw := cloudwatch.NewFromConfig(cfg)

	return scw, nil
}

func GetAWSCloudWatchLogsClientV1(accessKeyID string, secretAccessKey string, region string) (*cloudwatchlogsV1.CloudWatchLogs, error) {
	var sess *session.Session
	var err error
	if len(accessKeyID) == 0 || len(secretAccessKey) == 0 {
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region: aws.String(region),
			},
		})
	} else {
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region:      aws.String(region),
				Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
			},
		})
	}
	if err != nil {
		return nil, err
	}
	cwl := cloudwatchlogsV1.New(sess)
	return cwl, nil
}

package driver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"

	awssdkv1 "github.com/aws/aws-sdk-go/aws"
	awscredv1 "github.com/aws/aws-sdk-go/aws/credentials"
	awssessv1 "github.com/aws/aws-sdk-go/aws/session"
	cloudwatchlogsV1 "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func GetAWSCloudWatchClient(accessKeyID string, secretAccessKey string, region string) (*cloudwatch.Client, error) {
	var cfg aws.Config
	var err error

	if len(accessKeyID) == 0 || len(secretAccessKey) == 0 {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			return nil, err
		}
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")))
		if err != nil {
			return nil, err
		}
		cfg.Region = region

	}
	scw := cloudwatch.NewFromConfig(cfg)

	return scw, nil
}

func GetAWSCloudWatchLogsClientV1(accessKeyID string, secretAccessKey string, region string) (*cloudwatchlogsV1.CloudWatchLogs, error) {
	var sess *awssessv1.Session
	var err error
	if len(accessKeyID) == 0 || len(secretAccessKey) == 0 {
		sess, err = awssessv1.NewSessionWithOptions(awssessv1.Options{
			Config: awssdkv1.Config{
				Region: aws.String(region),
			},
		})
	} else {
		sess, err = awssessv1.NewSessionWithOptions(awssessv1.Options{
			Config: awssdkv1.Config{
				Region:      awssdkv1.String(region),
				Credentials: awscredv1.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
			},
		})
	}
	if err != nil {
		return nil, err
	}
	cwl := cloudwatchlogsV1.New(sess)
	return cwl, nil
}

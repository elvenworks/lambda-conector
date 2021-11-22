package driver

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	cloudwatchlogsV1 "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/elvenworks/lambda-conector/internal/domain"
)

type ILambda interface {
	GetAWSLambdaClient(config *domain.LambdaConfig) (*lambda.Client, error)
	GetAWSCloudwatchClient(config *domain.LambdaConfig) (*cloudwatch.Client, error)
	GetAWSCloudwatchLogsClient(config *domain.LambdaConfig) (*cloudwatchlogs.Client, error)
	GetAWSCloudWatchLogsClientV1(config *domain.LambdaConfig) (*cloudwatchlogsV1.CloudWatchLogs, error)
}

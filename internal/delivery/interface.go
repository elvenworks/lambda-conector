package delivery

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type ILambda interface {
	// CheckLambda(domain string) error
	GetAWSLambdaClient(config *LambdaConfig) (*lambda.Client, error)
	GetAWSCloudwatchClient(config *LambdaConfig) (*cloudwatch.Client, error)
	GetAWSCloudwatchLogsClient(config *LambdaConfig) (*cloudwatchlogs.Client, error)
	ConfigureAWSLambda(domain string, periodicidade int32, domainSettings map[string]string) (*LambdaConfig, error)
}

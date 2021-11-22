package delivery

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/elvenworks/lambda-conector/internal/domain"
)

type ILambda interface {
	// CheckLambda(domain string) error
	GetAWSLambdaClient(config *domain.LambdaConfig) (*lambda.Client, error)
	GetAWSCloudwatchClient(config *domain.LambdaConfig) (*cloudwatch.Client, error)
	GetAWSCloudwatchLogsClient(config *domain.LambdaConfig) (*cloudwatchlogs.Client, error)
}

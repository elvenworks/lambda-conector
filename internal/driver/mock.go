package driver

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	cloudwatchlogsV1 "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/elvenworks/lambda-conector/internal/domain"
	"github.com/stretchr/testify/mock"
)

type LambdaDeliveryMock struct {
	mock.Mock
}

func (m LambdaDeliveryMock) GetAWSLambdaClient(config *domain.LambdaConfig) (*lambda.Client, error) {
	args := m.Called(config)

	return args.Get(0).(*lambda.Client), args.Error(1)
}

func (m LambdaDeliveryMock) GetAWSCloudwatchClient(config *domain.LambdaConfig) (*cloudwatch.Client, error) {
	args := m.Called(config)

	return args.Get(0).(*cloudwatch.Client), args.Error(1)
}

func (m LambdaDeliveryMock) GetAWSCloudwatchLogsClient(config *domain.LambdaConfig) (*cloudwatchlogs.Client, error) {
	args := m.Called(config)

	return args.Get(0).(*cloudwatchlogs.Client), args.Error(1)
}

func (m LambdaDeliveryMock) GetAWSCloudwatchLogsV1Client(config *domain.LambdaConfig) (*cloudwatchlogsV1.CloudWatchLogs, error) {
	args := m.Called(config)

	return args.Get(0).(*cloudwatchlogsV1.CloudWatchLogs), args.Error(1)
}

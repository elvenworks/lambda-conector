package delivery

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
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

func (m LambdaDeliveryMock) GetAWSCloudwatchLogsClient(config *domain.LambdaConfig) (*cloudwatchlogs.CloudWatchLogs, error) {
	args := m.Called(config)

	return args.Get(0).(*cloudwatchlogs.CloudWatchLogs), args.Error(1)
}

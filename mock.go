package lambda

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/elvenworks/lambda-conector/domain"
	"github.com/stretchr/testify/mock"
)

type LambdaMock struct {
	mock.Mock
}

func (m LambdaMock) GetConfig() *domain.LambdaConfig {
	args := m.Called()
	return args.Get(0).(*domain.LambdaConfig)
}

func (m LambdaMock) GetLastLambdaRun() (*domain.LambdaLastRun, error) {
	args := m.Called()
	return args.Get(0).(*domain.LambdaLastRun), args.Error(1)
}

func (m LambdaMock) GetInvocationsAndErrors(startTime time.Time, endTime time.Time, id1 string, id2 string, period int32) (*cloudwatch.GetMetricDataOutput, error) {
	args := m.Called(startTime, endTime, id1, id2, period)
	return args.Get(0).(*cloudwatch.GetMetricDataOutput), args.Error(1)
}

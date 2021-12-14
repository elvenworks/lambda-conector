package lambda

import (
	"github.com/elvenworks/lambda-conector/internal/domain"
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

func (m LambdaMock) GetLogsLastErrorRun() (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

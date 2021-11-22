package lambda

import (
	"github.com/elvenworks/lambda-conector/internal/domain"
	"github.com/stretchr/testify/mock"
)

type LambdaMock struct {
	mock.Mock
}

func (m LambdaMock) GetLastLambdaRunMock(config domain.LambdaConfig) (*domain.LambdaLastRun, error) {
	args := m.Called(config)

	return args.Get(0).(*domain.LambdaLastRun), args.Error(1)
}

func (m LambdaMock) GetLogsLastErrorRunMock(config domain.LambdaConfig) (string, error) {
	args := m.Called(config)

	return "", args.Error(1)
}

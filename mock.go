package lambda

import (
	"github.com/elvenworks/lambda-conector/internal/domain"
	"github.com/stretchr/testify/mock"
)

type LambdaMock struct {
	mock.Mock
}

func (m LambdaMock) InitLambda(config domain.LambdaConfig) *Lambda {
	args := m.Called(config)

	return args.Get(0).(*Lambda)
}

func (m LambdaMock) GetConfig() *domain.LambdaConfig {
	lam := m.InitLambda(domain.LambdaConfig{})
	args := lam.GetConfig()

	return args
}

func (m LambdaMock) GetLastLambdaRunMock() (*domain.LambdaLastRun, error) {
	lam := m.InitLambda(domain.LambdaConfig{})
	args, err := lam.GetLastLambdaRun()

	return args, err
}

func (m LambdaMock) GetLogsLastErrorRunMock() (string, error) {
	lam := m.InitLambda(domain.LambdaConfig{})
	args, err := lam.GetLogsLastErrorRun()

	return args, err
}

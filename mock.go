package lambda

import "github.com/stretchr/testify/mock"

type LambdaMock struct {
	mock.Mock
}

func (m LambdaMock) GetLastLambdaRun(param LambdaParam) error {
	args := m.Called(param)

	return args.Error(0)
}

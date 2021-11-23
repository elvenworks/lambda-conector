package lambda

import "github.com/elvenworks/lambda-conector/internal/domain"

type ILambda interface {
	GetConfig() *domain.LambdaConfig
	GetLastLambdaRun() (*domain.LambdaLastRun, error)
	GetLogsLastErrorRun() (string, error)
}

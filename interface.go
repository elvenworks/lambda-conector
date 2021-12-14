package lambda

import "github.com/elvenworks/lambda-conector/domain"

type ILambda interface {
	GetConfig() *domain.LambdaConfig
	GetLastLambdaRun() (*domain.LambdaLastRun, error)
	GetLogsLastErrorRun() (string, error)
}

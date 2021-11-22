package lambda

import "github.com/elvenworks/lambda-conector/internal/domain"

type ILambda interface {
	// CheckLambda(domain string) error
	GetLastLambdaRun(config domain.LambdaConfig) (*domain.LambdaLastRun, error)
	GetLogsLastErrorRun(config domain.LambdaConfig) (string, error)
}

package lambda

type ILambda interface {
	// CheckLambda(domain string) error
	GetLastLambdaRun(param LambdaParam) error
}

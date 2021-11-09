package lambda

type ILambda interface {
	// CheckLambda(domain string) error
	GetLastLambdaRun(domain string) ([]byte, error)
}

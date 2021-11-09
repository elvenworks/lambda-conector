package delivery

import "github.com/aws/aws-sdk-go/service/lambda"

type ILambda interface {
	// CheckLambda(domain string) error
	GetAWSLambdaClient(config *LambdaConfig) (*lambda.Lambda, error)
}

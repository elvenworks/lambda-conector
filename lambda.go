package lambda

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/elvenworks/lambda-conector/internal/delivery"
)

func GetLastLambdaRun(domain string) (result []byte, err error) {

	config, err := delivery.ConfigureAWSLambda(domain)
	if err != nil {
		return nil, err
	}

	client, err := delivery.GetAWSLambdaClient(config)
	if err != nil {
		return nil, err
	}

	resultOutput, err := client.GetFunction(&lambda.GetFunctionInput{
		FunctionName: &config.FunctionName,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(resultOutput)
	// TODO: Implementar busca da ultima execução do lambda e o status

	return nil, nil
}

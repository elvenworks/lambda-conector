package delivery

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/elvenworks/lambda-conector/pkg/aws"
)

// // domain: account:password/region@functionName
func GetAWSLambdaClient(config *LambdaConfig) (*lambda.Lambda, error) {

	// confAWS, err := configureAWSLambda(domain)
	// if err != nil {
	// 	u.recordFailure(s, fmt.Sprintf("%v Lambda panic occured when split domain params: %v", strings.ToUpper(s.Method.String), err))
	// 	return nil, err
	// }

	// AWS API
	session := aws.NewSession(&aws.Config{
		AccessKeyID:     config.AccessKeyID,
		SecretAccessKey: config.SecretAccessKey,
		Region:          config.Region,
	})

	svc := lambda.New(session)

	return svc, nil
}

type LambdaConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	FunctionName    string
}

func ConfigureAWSLambda(domain string) (*LambdaConfig, error) {
	// Extract domain params
	splited := strings.LastIndex(domain, "/")
	if splited == 0 {
		return nil, errors.New("invalid domain format")
	}

	auth := strings.Split(domain[0:splited], ":")
	config := strings.Split(domain[splited+1:], "@")

	if len(auth) != 2 {
		return nil, errors.New("invalid authorization format")
	}

	if len(config) != 2 {
		return nil, errors.New("invalid region or function name")
	}

	region := config[0]
	functionName := config[1]

	return &LambdaConfig{
		AccessKeyID:     auth[0],
		SecretAccessKey: auth[1],
		Region:          region,
		FunctionName:    functionName,
	}, nil
}

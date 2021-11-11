package delivery

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
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

// // domain: account:password/region@functionName
func GetAWSCloudWatchClient(config *LambdaConfig) (*cloudwatch.CloudWatch, error) {

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

	scw := cloudwatch.New(session)

	return scw, nil
}

type LambdaConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Namespace       string
	FunctionName    string
	Period          int64
	MetricName      string
}

func ConfigureAWSLambda(domain string, periodicidade int64, domainSettings map[string]string) (*LambdaConfig, error) {
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

	if domainSettings["name_space"] == "" {
		return nil, errors.New("invalid name space")
	}

	if domainSettings["metric_name"] == "" {
		return nil, errors.New("invalid metric name")
	}

	return &LambdaConfig{
		AccessKeyID:     auth[0],
		SecretAccessKey: auth[1],
		Region:          region,
		FunctionName:    functionName,
		Period:          periodicidade,
		Namespace:       domainSettings["name_space"],
		MetricName:      domainSettings["metric_name"],
	}, nil
}

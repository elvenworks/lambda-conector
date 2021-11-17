package delivery

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

// // domain: account:password/region@functionName
func GetAWSLambdaClient(lambdaConfig *LambdaConfig) (*lambda.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("unable to load Lambda SDK config, %v", err)
		return nil, err
	}

	svc := lambda.NewFromConfig(cfg)

	return svc, nil
}

// // domain: account:password/region@functionName
func GetAWSCloudWatchClient(lambdaConfig *LambdaConfig) (*cloudwatch.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("unable to load Cloudwatch SDK config, %v", err)
		return nil, err
	}
	scw := cloudwatch.NewFromConfig(cfg)

	return scw, nil
}

// // domain: account:password/region@functionName
func GetAWSCloudWatchLogsClient(lambdaConfig *LambdaConfig) (*cloudwatchlogs.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("unable to load Cloudwatch Logs SDK config, %v", err)
		return nil, err
	}
	scl := cloudwatchlogs.NewFromConfig(cfg)

	return scl, nil
}

type LambdaConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Namespace       string
	FunctionName    string
	Period          int32
	MetricName      string
	Stat            string
}

func ConfigureAWSLambda(domain string, periodicidade int32) (*LambdaConfig, error) {
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
		Period:          periodicidade,
		Namespace:       "AWS/Lambda",
		MetricName:      "Errors",
		Stat:            "Average",
	}, nil
}

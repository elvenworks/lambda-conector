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
	"github.com/elvenworks/lambda-conector/internal/domain"
)

// // domain: account:password/region@functionName
func GetAWSLambdaClient(lambdaConfig *domain.LambdaConfig) (*lambda.Client, error) {

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
func GetAWSCloudWatchClient(lambdaConfig *domain.LambdaConfig) (*cloudwatch.Client, error) {

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
func GetAWSCloudWatchLogsClient(lambdaConfig *domain.LambdaConfig) (*cloudwatchlogs.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(lambdaConfig.Region),
	)
	if err != nil {
		log.Fatalf("unable to load Cloudwatch Logs SDK config, %v", err)
		return nil, err
	}
	scl := cloudwatchlogs.NewFromConfig(cfg)

	return scl, nil
}

func ConfigureAWSLambda(host string, period int32) (*domain.LambdaConfig, error) {
	// Extract host params
	splited := strings.LastIndex(host, "/")
	if splited == 0 {
		return nil, errors.New("invalid host format")
	}

	auth := strings.Split(host[0:splited], ":")
	config := strings.Split(host[splited+1:], "@")

	if len(auth) != 2 {
		return nil, errors.New("invalid authorization format")
	}

	if len(config) != 2 {
		return nil, errors.New("invalid region or function name")
	}

	region := config[0]
	functionName := config[1]

	return &domain.LambdaConfig{
		AccessKeyID:       auth[0],
		SecretAccessKey:   auth[1],
		Region:            region,
		FunctionName:      functionName,
		Period:            period,
		Namespace:         "AWS/Lambda",
		MetricErrors:      "Errors",
		MetricInvocations: "Invocations",
		Stat:              "Sum",
	}, nil

}

package delivery

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/elvenworks/lambda-conector/internal/domain"
)

func GetAWSLambdaClient(lambdaConfig *domain.LambdaConfig) (*lambda.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(lambdaConfig.Region),
	)
	if err != nil {
		log.Fatalf("unable to load Lambda SDK config, %v", err)
		return nil, err
	}

	svc := lambda.NewFromConfig(cfg)

	return svc, nil
}

func GetAWSCloudWatchClient(lambdaConfig *domain.LambdaConfig) (*cloudwatch.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(lambdaConfig.Region),
	)
	if err != nil {
		log.Fatalf("unable to load Cloudwatch SDK config, %v", err)
		return nil, err
	}
	scw := cloudwatch.NewFromConfig(cfg)

	return scw, nil
}

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

package lambda

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/elvenworks/lambda-conector/internal/delivery"
)

func GetLastLambdaRun(domain string, periodicidade int64, domainSettings map[string]string) (result []byte, err error) {

	config, err := delivery.ConfigureAWSLambda(domain, periodicidade, domainSettings)
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

	cliCW, err := delivery.GetAWSCloudWatchClient(config)
	if err != nil {
		return nil, err
	}

	constFuncName := "FunctionName"
	req, output := cliCW.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
		StartTime:  &time.Time{},
		EndTime:    &time.Time{},
		Namespace:  &config.Namespace,
		MetricName: &config.MetricName, //Invocations
		Period:     &periodicidade,
		Dimensions: []*cloudwatch.Dimension{&cloudwatch.Dimension{
			Name:  &constFuncName,
			Value: &config.FunctionName,
		}},
	})
	err = req.Send()
	if err != nil {
		fmt.Println(err)
	}

	output2, err := cliCW.GetMetricStatistics(&cloudwatch.GetMetricStatisticsInput{
		StartTime:  &time.Time{},
		EndTime:    &time.Time{},
		Namespace:  &config.Namespace,
		MetricName: &config.MetricName, //Invocations
		Period:     &periodicidade,
		Dimensions: []*cloudwatch.Dimension{&cloudwatch.Dimension{
			Name:  &constFuncName,
			Value: &config.FunctionName,
		}},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req)
	fmt.Println(output)
	fmt.Println(output2)

	return []byte("OK!"), nil
}

package lambda

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/elvenworks/lambda-conector/internal/delivery"
)

type LambdaParam struct {
	Domain string
	Period int32
}

func GetLastLambdaRun(lambdaParam LambdaParam) (err error) {

	config, err := delivery.ConfigureAWSLambda(lambdaParam.Domain, lambdaParam.Period)
	if err != nil {
		log.Fatalf("unable to get AWS config, %v", err)
		return err
	}
	// fmt.Println(config)
	// client, err := delivery.GetAWSLambdaClient(config)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(client)
	// resultOutput, err := client.GetFunction(context.TODO(), &lambda.GetFunctionInput{
	// 	FunctionName: &config.FunctionName,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(resultOutput)

	cliCW, err := delivery.GetAWSCloudWatchClient(config)
	if err != nil {
		log.Fatalf("unable to get cloudwatch client, %v", err)
		return err
	}
	// fmt.Println(cliCW)
	endTime := time.Now()
	startTime := time.Now().Add(time.Minute * 1 * -1)
	id := "e1"
	output2, err := cliCW.GetMetricData(context.TODO(), &cloudwatch.GetMetricDataInput{
		StartTime: &startTime,
		EndTime:   &endTime,
		MetricDataQueries: []types.MetricDataQuery{
			types.MetricDataQuery{
				Id: &id,
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						MetricName: &config.MetricName,
						Namespace:  &config.Namespace,
					},
					Period: &lambdaParam.Period,
					Stat:   &config.Stat,
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to get metric data, %v", err)
		return err
	}

	if len(output2.MetricDataResults[0].Values) == 0 {
		return nil
	}
	// fmt.Println(len(output2.MetricDataResults))
	lastTimestamp := output2.MetricDataResults[0].Timestamps[0]
	lastErr := output2.MetricDataResults[0].Values[0]
	fmt.Println(lastErr)
	if lastErr != 0 {
		cliCWL, err := delivery.GetAWSCloudWatchLogsClient(config)
		if err != nil {
			log.Fatalf("unable to get cloudwatchlogs client, %v", err)
			return err
		}

		logPointer := "ERROR"
		outputLog, err := cliCWL.GetLogRecord(context.TODO(), &cloudwatchlogs.GetLogRecordInput{
			LogRecordPointer: &logPointer,
		})
		if err != nil {
			log.Fatalf("unable to get cloudwatch logs, %v", err)
			return err
		}
		fmt.Println(outputLog)
	}
	fmt.Println(lastTimestamp)
	fmt.Println(lastErr)

	return nil
}

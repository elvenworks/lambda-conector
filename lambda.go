package lambda

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
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
	startTime := time.Now().Add(time.Hour * 24 * -1)
	id1, id2 := "e1", "e2"

	output2, err := cliCW.GetMetricData(context.TODO(), &cloudwatch.GetMetricDataInput{
		StartTime: &startTime,
		EndTime:   &endTime,
		MetricDataQueries: []types.MetricDataQuery{
			{
				Id: &id1,
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						MetricName: &config.MetricErrors,
						Namespace:  &config.Namespace,
					},
					Period: &lambdaParam.Period,
					Stat:   &config.Stat,
				},
			},
			{
				Id: &id2,
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						MetricName: &config.MetricInvocations,
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

		if len(output2.MetricDataResults[1].Timestamps) == 0 {

			return errors.New("the last run still results in error")
		} else {
			return nil
		}

	}
	// fmt.Println(len(output2.MetricDataResults))
	lastTimestamp := output2.MetricDataResults[0].Timestamps[0]
	lastErr := output2.MetricDataResults[0].Values[0]
	fmt.Println(lastErr)
	if lastErr != 0 {
		txtError := "timestamp: "
		return errors.New("the last run results in error, " + txtError + lastTimestamp.String())

		// cliCWL, err := delivery.GetAWSCloudWatchLogsClient(config)
		// if err != nil {
		// 	log.Fatalf("unable to get cloudwatchlogs client, %v", err)
		// 	return err
		// }

		// logPointer := "ERROR"
		// outputLog, err := cliCWL.GetLogRecord(context.TODO(), &cloudwatchlogs.GetLogRecordInput{
		// 	LogRecordPointer: &logPointer,
		// })
		// if err != nil {
		// 	log.Fatalf("unable to get cloudwatch logs, %v", err)
		// 	return err
		// }
		// fmt.Println(outputLog)
	}
	// fmt.Println(lastTimestamp)
	// fmt.Println(lastErr)

	return nil
}

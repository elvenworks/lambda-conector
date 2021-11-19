package lambda

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/elvenworks/lambda-conector/internal/delivery"
)

func initSessionV1(config delivery.LambdaConfig) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(config.Region),
			Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.SecretAccessKey, ""),
		},
	})
	if err != nil {
		panic(err)
	}
	cwl = cloudwatchlogs.New(sess)
}

var (
	ccw *cloudwatch.Client
	cwl *cloudwatchlogs.CloudWatchLogs
)

type LambdaParam struct {
	Domain       string
	Period       int32
	LogGroupName string
}

func GetLastLambdaRun(lambdaParam LambdaParam) (err error) {

	config, err := delivery.ConfigureAWSLambda(lambdaParam.Domain, lambdaParam.Period)
	if err != nil {
		log.Fatalf("unable to get AWS config, %v", err)
		return err
	}

	initSessionV1(*config)

	ccw, err = delivery.GetAWSCloudWatchClient(config)
	if err != nil {
		log.Fatalf("unable to get cloudwatch client, %v", err)
		return err
	}

	endTime := time.Now()
	startTime := time.Now().Add(time.Hour * 24 * -1)
	id1, id2 := "e1", "e2"

	output2, err := ccw.GetMetricData(context.TODO(), &cloudwatch.GetMetricDataInput{
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
			return errors.New("no invocations today")
		} else {
			return nil
		}
	}

	lastErr := output2.MetricDataResults[0].Values[0]
	if lastErr != 0 {
		message, err := getCWErrorLogMessage(lambdaParam.LogGroupName)
		if err != nil {
			log.Fatalf("unable to get cloudwatchlogs client, %v", err)
			return err
		}
		return errors.New("the last run results in error, " + message)
	}

	return nil
}

func getCWErrorLogMessage(logGroupName string) (string, error) {
	output, err := cwl.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: &logGroupName,
		Descending:   aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("unable to get cloudwatch logs, %v", err)
		return "unable to get cloudwatch logs", err
	}

	output2, err := cwl.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  &logGroupName,
		LogStreamName: output.LogStreams[0].LogStreamName,
	})
	if err != nil {
		log.Fatalf("unable to get cloudwatch logs, %v", err)
		return "unable to get cloudwatch logs", err
	}

	eventsSlice := output2.Events
	for i := 0; i < len(eventsSlice); i++ {
		if strings.Contains(*eventsSlice[i].Message, "ERROR") {
			return *eventsSlice[i].Message, nil
		}
	}
	return "error log not found", errors.New("error log not found")
}

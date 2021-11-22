package lambda

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go/aws"
	cloudwatchlogsV1 "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/elvenworks/lambda-conector/internal/domain"
	"github.com/elvenworks/lambda-conector/internal/driver"
)

func GetLastLambdaRun(config domain.LambdaConfig) (*domain.LambdaLastRun, error) {

	ccw, err := driver.GetAWSCloudWatchClient(&config)
	if err != nil {
		// log.Fatalf("unable to get cloudwatch client, %v", err)
		return nil, err
	}

	endTime := time.Now()
	startTime := time.Now().Add(time.Second * time.Duration(config.Period) * 2 * -1)
	id1, id2 := "e1", "e2"

	output, err := ccw.GetMetricData(context.TODO(), &cloudwatch.GetMetricDataInput{
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
					Period: &config.Period,
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
					Period: &config.Period,
					Stat:   &config.Stat,
				},
			},
		},
	})
	if err != nil {
		// log.Fatalf("unable to get metric data, %v", err)
		return nil, err
	}

	if len(output.MetricDataResults[0].Values) == 0 {
		return nil, errors.New("no invocations for the period")
	}

	return &domain.LambdaLastRun{
		Timestamp:  output.MetricDataResults[0].Timestamps[0],
		ErrorCount: output.MetricDataResults[0].Values[0],
	}, nil
}

func GetLogsLastErrorRun(config domain.LambdaConfig) (string, error) {
	cwl, err := driver.GetAWSCloudWatchLogsClientV1(&config)
	if err != nil {
		// log.Fatalf("unable to get cloudwatch client, %v", err)
		return "", err
	}

	output, err := cwl.DescribeLogStreams(&cloudwatchlogsV1.DescribeLogStreamsInput{
		LogGroupName: &config.LogGroupName,
		Descending:   aws.Bool(true),
	})
	if err != nil {
		// log.Fatalf("unable to get cloudwatch logs streams, %v", err)
		return "", err
	}

	output2, err := cwl.GetLogEvents(&cloudwatchlogsV1.GetLogEventsInput{
		LogGroupName:  &config.LogGroupName,
		LogStreamName: output.LogStreams[0].LogStreamName,
	})
	if err != nil {
		// log.Fatalf("unable to get cloudwatch logs, %v", err)
		return "", err
	}

	eventsSlice := output2.Events
	for i := 0; i < len(eventsSlice); i++ {
		if strings.Contains(*eventsSlice[i].Message, "ERROR") {
			return *eventsSlice[i].Message, nil
		}
	}
	return "", errors.New("error log not found")

}

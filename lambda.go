package lambda

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	awssdkv1 "github.com/aws/aws-sdk-go/aws"
	cloudwatchlogsV1 "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/elvenworks/lambda-conector/domain"
	"github.com/elvenworks/lambda-conector/internal/driver"
	"github.com/sirupsen/logrus"
)

type Lambda struct {
	Clients domain.Clients
	config  domain.LambdaConfig
}

type InitConfig struct {
	AccessKeyID      string
	SecretAccessKey  string
	Region           string
	FunctionName     string
	Period           int32
	LogGroupName     string
	FlagSearchPeriod bool
}

func InitLambda(config InitConfig) *Lambda {

	ccw, err := driver.GetAWSCloudWatchClient(config.AccessKeyID, config.SecretAccessKey, config.Region)
	if err != nil {
		logrus.Error("unable to get cloudwatch client, %v", err)
	}

	ccwlv1, err := driver.GetAWSCloudWatchLogsClientV1(config.AccessKeyID, config.SecretAccessKey, config.Region)
	if err != nil {
		logrus.Error("unable to get cloudwatchlogs v1 client, %v", err)
	}

	return &Lambda{
		Clients: domain.Clients{
			Ccw:    *ccw,
			Ccwlv1: *ccwlv1,
		},
		config: domain.LambdaConfig{
			AccessKeyID:       config.AccessKeyID,
			SecretAccessKey:   config.SecretAccessKey,
			Region:            config.Region,
			FunctionName:      config.FunctionName,
			Period:            config.Period,
			LogGroupName:      config.LogGroupName,
			Namespace:         "AWS/Lambda",
			MetricErrors:      "Errors",
			MetricInvocations: "Invocations",
			Stat:              "Sum",
			FlagSearchPeriod:  config.FlagSearchPeriod,
		},
	}
}

func (l *Lambda) GetConfig() *domain.LambdaConfig {
	return &l.config
}

func (l *Lambda) GetLastLambdaRun() (*domain.LambdaLastRun, error) {

	endTime := time.Now()
	startTime := time.Now().Add(time.Second * time.Duration(l.GetConfig().Period) * 2 * -1)
	id1, id2 := "e1", "e2"

	output, err := l.Clients.Ccw.GetMetricData(context.TODO(), &cloudwatch.GetMetricDataInput{
		StartTime: &startTime,
		EndTime:   &endTime,
		MetricDataQueries: []types.MetricDataQuery{
			{
				Id: &id1,
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						MetricName: &l.GetConfig().MetricErrors,
						Namespace:  &l.GetConfig().Namespace,
					},
					Period: &l.GetConfig().Period,
					Stat:   &l.GetConfig().Stat,
				},
			},
			{
				Id: &id2,
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						MetricName: &l.GetConfig().MetricInvocations,
						Namespace:  &l.GetConfig().Namespace,
					},
					Period: &l.GetConfig().Period,
					Stat:   &l.GetConfig().Stat,
				},
			},
		},
	})
	if err != nil {
		logrus.Error("unable to get metric data, %v", err)
		return nil, err
	}

	if len(output.MetricDataResults[0].Values) == 0 {
		if l.GetConfig().FlagSearchPeriod {
			return nil, errors.New("no invocations for the period")
		} else {
			endTime = time.Now()
			startTime = time.Now().Add(time.Hour * 24 * -1)
			id1, id2 = "i1", "i2"
			output, err = l.Clients.Ccw.GetMetricData(context.TODO(), &cloudwatch.GetMetricDataInput{
				StartTime: &startTime,
				EndTime:   &endTime,
				MetricDataQueries: []types.MetricDataQuery{
					{
						Id: &id1,
						MetricStat: &types.MetricStat{
							Metric: &types.Metric{
								MetricName: &l.GetConfig().MetricErrors,
								Namespace:  &l.GetConfig().Namespace,
							},
							Period: &l.GetConfig().Period,
							Stat:   &l.GetConfig().Stat,
						},
					}},
			})
			if err != nil {
				logrus.Error("unable to get metric data, %v", err)
				return nil, err
			}

			return &domain.LambdaLastRun{
				Timestamp:  output.MetricDataResults[0].Timestamps[0],
				ErrorCount: output.MetricDataResults[0].Values[0],
			}, nil
		}
	}

	return &domain.LambdaLastRun{
		Timestamp:  output.MetricDataResults[0].Timestamps[0],
		ErrorCount: output.MetricDataResults[0].Values[0],
	}, nil
}

func (l *Lambda) GetLogsLastErrorRun() (string, error) {

	output, err := l.Clients.Ccwlv1.DescribeLogStreams(&cloudwatchlogsV1.DescribeLogStreamsInput{
		LogGroupName: &l.GetConfig().LogGroupName,
		Descending:   awssdkv1.Bool(true),
		OrderBy:      awssdkv1.String("LastEventTime"),
	})
	if err != nil {
		logrus.Error("unable to get cloudwatch logs streams, %v", err)
		return "", err
	}

	output2, err := l.Clients.Ccwlv1.GetLogEvents(&cloudwatchlogsV1.GetLogEventsInput{
		LogGroupName:  &l.GetConfig().LogGroupName,
		LogStreamName: output.LogStreams[0].LogStreamName,
	})
	if err != nil {
		logrus.Error("unable to get cloudwatch logs, %v", err)
		return "", err
	}

	eventsSlice := output2.Events
	for i := 0; i < len(eventsSlice); i++ {
		if strings.Contains(*eventsSlice[i].Message, "ERROR") {
			return *eventsSlice[i].Message, nil
		}
	}
	return "", nil

}

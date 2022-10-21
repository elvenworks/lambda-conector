package lambda

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
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

	return &Lambda{
		Clients: domain.Clients{
			Ccw: *ccw,
		},
		config: domain.LambdaConfig{
			AccessKeyID:       config.AccessKeyID,
			SecretAccessKey:   config.SecretAccessKey,
			Region:            config.Region,
			FunctionName:      config.FunctionName,
			Period:            config.Period,
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

	if err := validatePeriod(l.GetConfig().Period); err != nil {
		return nil, err
	}

	startTime := time.Now().Add(time.Second * time.Duration(l.GetConfig().Period) * 2 * -1)
	endTime := time.Now()

	id1, id2 := "inv_per", "err_per"
	output, err := l.GetInvocationsAndErrors(startTime, endTime, id1, id2, l.config.Period)

	if err != nil {
		logrus.Error("unable to get metric data, %v", err)
		return nil, err
	}

	// no invocations occurred in the period so check FlagSearchPeriod or the last 24 hours
	if len(output.MetricDataResults[0].Values) == 0 {
		if l.GetConfig().FlagSearchPeriod {
			return &domain.LambdaLastRun{
				ErrorCount: 1,
				Message:    "no invocations for the period",
			}, nil
		}
		startTime = time.Now().Add(time.Hour * 24 * -1)
		endTime = time.Now()
		id1, id2 = "inv_l24", "err_l24"

		//for last 24h invocations/errors will be grouped every minute
		output, err = l.GetInvocationsAndErrors(startTime, endTime, id1, id2, int32(60))

		if err != nil {
			logrus.Error("unable to get metric data last 24 hours, %v", err)
			return nil, err
		}

	}

	// Some invocations and Some error occurred in the period
	// position 0: sum of invocations
	// position 1: sum of errors
	if len(output.MetricDataResults[0].Values) != 0 && len(output.MetricDataResults[1].Values) != 0 {
		return &domain.LambdaLastRun{
			Timestamp:  output.MetricDataResults[0].Timestamps[0],
			ErrorCount: output.MetricDataResults[1].Values[0],
			Message:    fmt.Sprintf("the last %v invocations had %v errors", output.MetricDataResults[0].Values[0], output.MetricDataResults[1].Values[0]),
		}, nil
	}

	// no invocations occurred the last 24 hours
	return &domain.LambdaLastRun{
		ErrorCount: -1,
		Message:    "no invocations occurred the last 24 hours",
	}, nil

}

func (l *Lambda) GetInvocationsAndErrors(startTime time.Time, endTime time.Time, id1 string, id2 string, period int32) (*cloudwatch.GetMetricDataOutput, error) {

	output, err := l.Clients.Ccw.GetMetricData(context.TODO(), &cloudwatch.GetMetricDataInput{
		StartTime: &startTime,
		EndTime:   &endTime,
		MetricDataQueries: []types.MetricDataQuery{
			{
				Id: &id1,
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						MetricName: &l.GetConfig().MetricInvocations,
						Namespace:  &l.GetConfig().Namespace,
					},
					Period: &period,
					Stat:   &l.GetConfig().Stat,
				},
			},
			{
				Id: &id2,
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						MetricName: &l.GetConfig().MetricErrors,
						Namespace:  &l.GetConfig().Namespace,
					},
					Period: &period,
					Stat:   &l.GetConfig().Stat,
				},
			},
		},
	})

	return output, err
}

func validatePeriod(period int32) error {
	if period < 60 {
		for _, p := range []int32{1, 5, 10, 30} {
			if period == p {
				return nil
			}
		}
	}
	if period%60 == 0 {
		return nil
	}
	return errors.New("period must be a value in the set [ 1, 5, 10, 30 ] or any multiple of 60")
}

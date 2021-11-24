package domain

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	cloudwatchlogsV1 "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type LambdaConfig struct {
	AccessKeyID       string
	SecretAccessKey   string
	Region            string
	Namespace         string
	FunctionName      string
	Period            int32
	LogGroupName      string
	MetricErrors      string
	MetricInvocations string
	Stat              string
}

type LambdaLastRun struct {
	Timestamp  time.Time
	ErrorCount float64
}

type Clients struct {
	Cl     lambda.Client
	Ccw    cloudwatch.Client
	Ccwl   cloudwatchlogs.Client
	Ccwlv1 cloudwatchlogsV1.CloudWatchLogs
}

package domain

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type LambdaConfig struct {
	AccessKeyID       string
	SecretAccessKey   string
	Region            string
	Namespace         string
	FunctionName      string
	Period            int32
	MetricErrors      string
	MetricInvocations string
	Stat              string
	DimensionName     string
	FlagSearchPeriod  bool
}

type LambdaLastRun struct {
	Timestamp  time.Time
	ErrorCount float64
	Message    string
}

type Clients struct {
	Ccw cloudwatch.Client
}

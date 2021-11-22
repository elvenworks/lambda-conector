package domain

import "time"

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

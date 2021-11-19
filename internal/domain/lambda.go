package domain

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
}

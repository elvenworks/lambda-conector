package driver

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	cloudwatchlogsV1 "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/elvenworks/lambda-conector/internal/domain"
)

func TestGetAWSLambdaClient(t *testing.T) {
	type args struct {
		lambdaConfig *domain.LambdaConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *lambda.Client
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				lambdaConfig: &domain.LambdaConfig{
					Region:            "us-east-1",
					FunctionName:      "stop_start_rds_instance",
					Namespace:         "AWS/Lambda",
					LogGroupName:      "/aws/lambda/stop_start_rds_instance",
					MetricErrors:      "Errors",
					MetricInvocations: "Invocations",
					Stat:              "Sum",
					Period:            60,
				},
			},
			want:    &lambda.Client{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSLambdaClient(tt.args.lambdaConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSLambdaClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("GetAWSLambdaClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAWSCloudWatchClient(t *testing.T) {
	type args struct {
		lambdaConfig *domain.LambdaConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *cloudwatch.Client
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				lambdaConfig: &domain.LambdaConfig{
					Region:            "us-east-1",
					FunctionName:      "stop_start_rds_instance",
					Namespace:         "AWS/Lambda",
					LogGroupName:      "/aws/lambda/stop_start_rds_instance",
					MetricErrors:      "Errors",
					MetricInvocations: "Invocations",
					Stat:              "Sum",
					Period:            60,
				},
			},
			want:    &cloudwatch.Client{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSCloudWatchClient(tt.args.lambdaConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSCloudWatchClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("GetAWSLambdaClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAWSCloudWatchLogsClient(t *testing.T) {
	type args struct {
		lambdaConfig *domain.LambdaConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *cloudwatchlogs.Client
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				lambdaConfig: &domain.LambdaConfig{
					Region:            "us-east-1",
					FunctionName:      "stop_start_rds_instance",
					Namespace:         "AWS/Lambda",
					LogGroupName:      "/aws/lambda/stop_start_rds_instance",
					MetricErrors:      "Errors",
					MetricInvocations: "Invocations",
					Stat:              "Sum",
					Period:            60,
				},
			},
			want:    &cloudwatchlogs.Client{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSCloudWatchLogsClient(tt.args.lambdaConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSCloudWatchLogsClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("GetAWSLambdaClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAWSCloudWatchLogsClientV1(t *testing.T) {
	type args struct {
		config *domain.LambdaConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *cloudwatchlogsV1.CloudWatchLogs
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				config: &domain.LambdaConfig{
					Region:            "us-east-1",
					FunctionName:      "stop_start_rds_instance",
					Namespace:         "AWS/Lambda",
					LogGroupName:      "/aws/lambda/stop_start_rds_instance",
					MetricErrors:      "Errors",
					MetricInvocations: "Invocations",
					Stat:              "Sum",
					Period:            60,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSCloudWatchLogsClientV1(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSCloudWatchLogsClientV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("GetAWSLambdaClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

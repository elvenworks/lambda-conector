package delivery

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
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
		// 		{
		// 			name: "Success",
		// 			args: args{
		// 				lambdaConfig: &domain.LambdaConfig{},
		// 				// 	AccessKeyID:       "account",
		// 				// 	SecretAccessKey:   "password",
		// 				// 	Region:            "region",
		// 				// 	FunctionName:      "functionName",
		// 				// 	Namespace:         "AWS/Lambda",
		// 				// 	MetricErrors:      "Errors",
		// 				// 	MetricInvocations: "Invocations",
		// 				// 	Stat:              "Sum",
		// 				// 	Period:            60,
		// 				// },
		// 			},
		// 			want:    LambdaDeliveryMock{},
		// 			wantErr: false,
		// 		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSLambdaClient(tt.args.lambdaConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSLambdaClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSCloudWatchClient(tt.args.lambdaConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSCloudWatchClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAWSCloudWatchClient() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSCloudWatchLogsClient(tt.args.lambdaConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSCloudWatchLogsClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAWSCloudWatchLogsClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

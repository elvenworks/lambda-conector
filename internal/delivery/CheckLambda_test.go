package delivery

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func TestGetAWSLambdaClient(t *testing.T) {
	type args struct {
		config *LambdaConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *lambda.Lambda
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSLambdaClient(tt.args.config)
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
		config *LambdaConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *cloudwatch.CloudWatch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSCloudWatchClient(tt.args.config)
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

func TestConfigureAWSLambda(t *testing.T) {
	type args struct {
		domain         string
		periodicidade  int64
		domainSettings map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *LambdaConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfigureAWSLambda(tt.args.domain, tt.args.periodicidade, tt.args.domainSettings)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigureAWSLambda() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigureAWSLambda() = %v, want %v", got, tt.want)
			}
		})
	}
}

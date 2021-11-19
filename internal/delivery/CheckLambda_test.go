package delivery

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/elvenworks/lambda-conector/internal/domain"
)

func TestConfigureAWSLambda(t *testing.T) {
	type args struct {
		domain        string
		periodicidade int32
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.LambdaConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfigureAWSLambda(tt.args.domain, tt.args.periodicidade)
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
		// TODO: Add test cases.
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

package lambda

import (
	"reflect"
	"testing"

	"github.com/elvenworks/lambda-conector/internal/domain"
)

func TestGetLastLambdaRun(t *testing.T) {
	type args struct {
		config domain.LambdaConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.LambdaLastRun
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLastLambdaRun(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastLambdaRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastLambdaRun() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLogsLastErrorRun(t *testing.T) {
	type args struct {
		config domain.LambdaConfig
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLogsLastErrorRun(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLogsLastErrorRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLogsLastErrorRun() = %v, want %v", got, tt.want)
			}
		})
	}
}

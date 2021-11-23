package lambda

import (
	"reflect"
	"testing"

	"github.com/elvenworks/lambda-conector/internal/domain"
)

func TestInitLambda(t *testing.T) {
	type args struct {
		config domain.LambdaConfig
	}
	tests := []struct {
		name string
		args args
		want *LambdaMock
	}{
		{
			name: "Success init",
			args: args{
				config: domain.LambdaConfig{},
			},
			want: &LambdaMock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitLambda(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitLambda() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLambda_GetConfig(t *testing.T) {
	type fields struct {
		Clients domain.Clients
		config  domain.LambdaConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.LambdaConfig
	}{
		{
			name: "Success",
			fields: fields{
				config:  domain.LambdaConfig{},
				Clients: domain.Clients{},
			},
			want: &domain.LambdaConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lambda{
				Clients: tt.fields.Clients,
				config:  tt.fields.config,
			}
			if got := l.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lambda.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLambda_GetLastLambdaRun(t *testing.T) {
	type fields struct {
		Clients domain.Clients
		config  domain.LambdaConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    *domain.LambdaLastRun
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				config:  domain.LambdaConfig{},
				Clients: domain.Clients{},
			},
			want:    &domain.LambdaLastRun{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lambda{
				Clients: tt.fields.Clients,
				config:  tt.fields.config,
			}
			got, err := l.GetLastLambdaRun()
			if (err != nil) != tt.wantErr {
				t.Errorf("Lambda.GetLastLambdaRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lambda.GetLastLambdaRun() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLambda_GetLogsLastErrorRun(t *testing.T) {
	type fields struct {
		Clients domain.Clients
		config  domain.LambdaConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				config:  domain.LambdaConfig{},
				Clients: domain.Clients{},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lambda{
				Clients: tt.fields.Clients,
				config:  tt.fields.config,
			}
			got, err := l.GetLogsLastErrorRun()
			if (err != nil) != tt.wantErr {
				t.Errorf("Lambda.GetLogsLastErrorRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lambda.GetLogsLastErrorRun() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestGetLastLambdaRun(t *testing.T) {
// 	type args struct {
// 		config domain.LambdaConfig
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    float64
// 		wantErr bool
// 	}{
// 		{
// 			name: "Success LOGS",
// 			args: args{
// 				config: domain.LambdaConfig{
// 					Region:            "us-east-1",
// 					FunctionName:      "stop_start_rds_instance",
// 					Namespace:         "AWS/Lambda",
// 					LogGroupName:      "/aws/lambda/stop_start_rds_instance",
// 					MetricErrors:      "Errors",
// 					MetricInvocations: "Invocations",
// 					Stat:              "Sum",
// 					Period:            86400,
// 				},
// 			},
// 			want:    0,
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetLastLambdaRun(tt.args.config)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetLastLambdaRun() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got.ErrorCount, tt.want) {
// 				t.Errorf("GetLastLambdaRun() = %v, want %v", got.ErrorCount, tt.want)
// 			}
// 		})
// 	}
// }

// func TestGetLastLambdaRun2(t *testing.T) {
// 	type args struct {
// 		config domain.LambdaConfig
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *domain.LambdaLastRun
// 		wantErr bool
// 	}{
// 		{
// 			name: "Success No invocations",
// 			args: args{
// 				config: domain.LambdaConfig{
// 					Region:            "us-east-1",
// 					FunctionName:      "stop_start_rds_instance",
// 					Namespace:         "AWS/Lambda",
// 					LogGroupName:      "/aws/lambda/stop_start_rds_instance",
// 					MetricErrors:      "Errors",
// 					MetricInvocations: "Invocations",
// 					Stat:              "Sum",
// 					Period:            60,
// 				},
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetLastLambdaRun(tt.args.config)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetLastLambdaRun() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetLastLambdaRun() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestGetLogsLastErrorRun(t *testing.T) {
// 	type args struct {
// 		config domain.LambdaConfig
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		{
// 			name: "Success",
// 			args: args{
// 				config: domain.LambdaConfig{
// 					Region:            "us-east-1",
// 					FunctionName:      "stop_start_rds_instance",
// 					Namespace:         "AWS/Lambda",
// 					LogGroupName:      "/aws/lambda/stop_start_rds_instance",
// 					MetricErrors:      "Errors",
// 					MetricInvocations: "Invocations",
// 					Stat:              "Sum",
// 					Period:            60,
// 				},
// 			},
// 			want:    "ERROR",
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetLogsLastErrorRun(tt.args.config)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetLogsLastErrorRun() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !strings.Contains(got, tt.want) {
// 				t.Errorf("GetLogsLastErrorRun() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

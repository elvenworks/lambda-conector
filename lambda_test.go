package lambda

import (
	"reflect"
	"testing"
)

func TestGetLastLambdaRun(t *testing.T) {
	type args struct {
		lambdaParam LambdaParam
	}
	tests := []struct {
		name       string
		args       args
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetLastLambdaRun(tt.args.lambdaParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastLambdaRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetLastLambdaRun() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
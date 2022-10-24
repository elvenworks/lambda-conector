package lambda

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/elvenworks/lambda-conector/domain"
)

type ILambda interface {
	GetConfig() *domain.LambdaConfig
	GetLastLambdaRun() (*domain.LambdaLastRun, error)
	GetInvocationsAndErrors(startTime time.Time, endTime time.Time, id1 string, id2 string, period int32) (*cloudwatch.GetMetricDataOutput, error)
}

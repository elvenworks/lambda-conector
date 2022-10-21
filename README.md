# Simple connector for AWS Lambda with Cloudwatch

 [![SDK Documentation](https://img.shields.io/badge/SDK-Documentation-blue)](https://aws.github.io/aws-sdk-go-v2/docs/) [![API Reference](https://img.shields.io/badge/API-reference-blue.svg)](https://pkg.go.dev/mod/github.com/aws/aws-sdk-go-v2) [![CloudWatch Documentation](https://img.shields.io/badge/CloudWatch-Documentation-blue)](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/cloudwatch)


`lambda-conector` is a simple client AWS Lambda with CloudWatch for check invocations and error occurred in a time interval.

This version requires a minimum version of `Go 1.15`.

## Getting started
To get started working with this conector setup your project for Go modules, and retrieve this project's dependencies with `go get`.

###### Initialize Project
```sh
$ mkdir hello-lambda-conector
$ cd hello-lambda-conector
$ go mod init hello-lambda-conector
```
###### Add lambda-conector Dependencies

```sh
$ go get github.com/elvenworks/lambda-conector
```

## Operations available for this version

What operations can this connector perform?
* Fetches the sum of invocations and errors of a function for a given period. If the informed period does not have invocations, a new check of the last 24 hours will be carried out to obtain the last calls

### Usage - GetLastLambdaRun 
Return `'LambdaLastRun' | 'error' `: 
```go
type LambdaLastRun struct {
    Timestamp  time.Time
    ErrorCount float64
    Message    string
}
```

###### Write Code
In your preferred editor add the following content to main.go

```go
package main

import (
	lambdaconector "github.com/elvenworks/lambda-conector"
	"github.com/sirupsen/logrus"
)

func main() {

    // Initialize the AWS SDK with or Settings
    config := lambdaconector.InitConfig{
		AccessKeyID:     "",
		SecretAccessKey: "",
		Region:          "",
		FunctionName:    "",
		Period:          60,
		FlagSearchPeriod: false,
	}
	lamdba := lambdaconector.InitLambda(config)

	lambdaLastRun, err := lamdba.GetLastLambdaRun()
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("last execution: ", lambdaLastRun.Timestamp)
	logrus.Info("message: ", lambdaLastRun.Message)
	logrus.Info("errors: ", lambdaLastRun.ErrorCount)

}
```
###### Compile and Execute
```sh
$ go run .
INFO[0000] last execution: 2022-10-21 18:27:00 +0000 UTC 
INFO[0000] message: the last 10 invocations had 0 errors for the period 
INFO[0000] errors: 0 
```

package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/enkhalifapro/users-serverless/internal/users/api"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {
	// Setup logger.
	logger, _ := zap.NewProduction()

	h := api.BuildHTTPHandler(logger)
	adapter := httpadapter.New(h)

	logrus.Debug("Starting Lambda")
	lambda.Start(adapter.Proxy)
}

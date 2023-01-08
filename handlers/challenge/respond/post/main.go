package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialise zap logger")
	}
	log = logger
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	return
}

func main() {
	lambda.Start(Handler)
}

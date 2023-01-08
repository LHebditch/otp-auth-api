package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/LHebditch/auth/db"
	"github.com/LHebditch/auth/handlers/challenge"
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

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	challengeStore, err := db.NewChallengStore(ctx)
	if err != nil {
		panic("could not create challenge store: " + err.Error())
	}
	challenger := challenge.NewChallenger(log, challengeStore, os.Getenv("SENDER_EMAIL"), os.Getenv("SMS_SENDER_ID"))
	id, err := challenger.SendChallenge(ctx, GetSubject(event.Body))
	if err != nil {
		log.Error("failed to send challenge: " + err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}
	response, err := json.Marshal(struct{ id string }{
		id: id,
	})
	if err != nil {
		log.Error("failed to parse id into response: " + err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(response),
	}, nil
}

func GetSubject(eventBody string) string {
	subj := struct{ subject string }{}
	err := json.Unmarshal([]byte(eventBody), &subj)
	if err != nil {
		panic("cannot extract subject from event body")
	}
	return subj.subject
}

func main() {
	lambda.Start(Handler)
}

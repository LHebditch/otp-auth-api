package db

import (
	"context"
	"os"

	"github.com/LHebditch/auth/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

type IDynamoClient interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) 
}

type ChallengeStore struct {
	Client IDynamoClient
}

func NewChallengStore(ctx context.Context) (ChallengeStore, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return ChallengeStore{}, err
	}	
	client := dynamodb.NewFromConfig(cfg)
	return ChallengeStore{
		Client: client,
	}, nil
}

func (c ChallengeStore) SaveOTP(ctx context.Context, subject, otp string) (id string, err error) {
	id = uuid.New().String()
	entry := models.NewChallengeEntry(id, subject, otp)
	item, err := attributevalue.MarshalMap(entry)
	if err != nil {
		return
	}
	// TODO what do we do about the _pk and _sk values?
	c.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("CHALLENGE_TABLE_NAME")),
		Item: item,
	})
	return
}
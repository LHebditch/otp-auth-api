package sendsms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

func Send(ctx context.Context, subject, sender string, otp string) error {
	cfg, err := config.LoadDefaultConfig(ctx) //Get local credentias ~/.aws/credentials
	if err != nil {
		return err
	}

	client := sns.NewFromConfig(cfg)
	params := &sns.PublishInput{
		Message:     aws.String("Here's your one time password: " + otp),
		PhoneNumber: aws.String(subject),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {StringValue: aws.String(sender), DataType: aws.String("String")},
		},
	}
	_, err = client.Publish(ctx, params)
	return err
}

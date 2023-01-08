package sendses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func Send(ctx context.Context, subject, otp, sender string) error {
	cfg, err := config.LoadDefaultConfig(ctx) //Get local credentias ~/.aws/credentials
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client := ses.NewFromConfig(cfg)
	_, err = client.SendEmail(ctx, &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{subject},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("Here's your one time password: " + otp),
				},
			},
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Web Auth UK: Please verify your login"),
			},
		},
		Source: aws.String(sender),
	})
	return err
}

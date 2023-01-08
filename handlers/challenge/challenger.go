package challenge

import (
	"context"
	"errors"
	"math/rand"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/LHebditch/auth/db"
	sendses "github.com/LHebditch/auth/services/ses"
	sendsms "github.com/LHebditch/auth/services/sns"
	"go.uber.org/zap"
)

type OTPChallenger interface {
	GenerateOTP() string
	SendChallenge(ctx context.Context, subject string) (id string, err error)
}

type SendChallenge func(ctx context.Context, subject, otp, sender string) error
type Challenger struct {
	log       *zap.Logger
	SendEmail SendChallenge
	SendSMS   SendChallenge
	SaveOTP   func(ctx context.Context, subject, otp string) (id string, err error)
	senderEmail string
	smsSenderId string
}

func (c Challenger) GenerateOTP() string {
	var sb strings.Builder
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	for i := 0; i < 6; i++ {
		num := r.Intn(10)
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}

func (c Challenger) SendChallenge(ctx context.Context, subject string) (id string, err error) {
	otp := c.GenerateOTP()
	// decide whether or not to sen email or sms
	if _, err := mail.ParseAddress(subject); err != nil {
		err = c.SendEmail(ctx, subject, otp, c.senderEmail)
	} else {
		// TODO - understand costs and look int SMS OTPs
		// err = c.SendSMS(ctx, subject, otp, c.smsSenderId)
		return "", errors.New("Invalid email. OTP was not sent")
	}
	if err != nil {
		return
	}
	id, err = c.SaveOTP(ctx, subject, otp)
	return
}

func NewChallenger(log *zap.Logger, challengeStore db.ChallengeStore, senderEmail, smsSenderID string) OTPChallenger {
	return Challenger{
		log:       log,
		SendEmail: sendses.Send,
		SendSMS:   sendsms.Send,
		SaveOTP:   challengeStore.SaveOTP,
		senderEmail: senderEmail,
		smsSenderId: smsSenderID,
	}
}

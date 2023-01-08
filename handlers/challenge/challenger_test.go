package challenge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestOTPGeneration(t *testing.T) {
	log, _ := zap.NewProduction()
	challenger := Challenger{
		log: log,
	}
	otp := challenger.GenerateOTP()

	fmt.Println(otp)
	assert.Equal(t, 6, len(otp))
}
package main

import (
	"context"

	sendsms "github.com/LHebditch/auth/services/sns"
)

func main(){
	sendsms.Send(context.TODO(), "+447917860887", "123456")
}
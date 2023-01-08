package models

import (
	"fmt"
	"time"
)

type ChallengeEntry struct {
	_pk string
	_sk string
	_gsi string
	ID string
	Subject string
	Code string
	Expiry time.Time
	Redeemed bool
	Attempts int
}

func NewChallengeEntry(id, subject,code string) ChallengeEntry {
	return ChallengeEntry {
		_pk: fmt.Sprintf("challenge/%s", id),
		_sk: "",
		_gsi1: "subject/" + subject + "/challenge",
		ID: id,
		Subject: subject,
		Code: code,
		Expiry: time.Now().Add(time.Minute * 10),
		Redeemed: false,
		Attempts: 0,
	}
}
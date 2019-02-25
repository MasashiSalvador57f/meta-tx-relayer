package services

import (
	"github.com/satori/go.uuid"
	"time"
)

type OneTimeMessageHandler struct {}

func NewOneTimeMessageHandler() OneTimeMessageIssuer {
	return new(OneTimeMessageHandler)
}

func (ot *OneTimeMessageHandler) IssueMsg(timeToExpiry time.Duration, seed string) string {
	uuid := uuid.NewV4()
	return uuid.String()
}


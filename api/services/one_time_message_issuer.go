package services

import (
	"crypto/sha256"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type OneTimeMessageHandler struct{}

func NewOneTimeMessageHandler() OneTimeMessageIssuer {
	return new(OneTimeMessageHandler)
}

func (ot *OneTimeMessageHandler) IssueMsg(timeToExpiry time.Duration, seed string) string {
	uuid := uuid.NewV4()
	sum := sha256.Sum256([]byte(uuid.String()))
	return fmt.Sprintf("%x", sum)
}

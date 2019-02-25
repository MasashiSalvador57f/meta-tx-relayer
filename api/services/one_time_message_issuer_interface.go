package services

import "time"

// OneTimeMessageIssuer is ...
type OneTimeMessageIssuer interface {
	IssueMsg(duration time.Duration, seed string) string
}


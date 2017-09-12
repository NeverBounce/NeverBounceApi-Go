package neverbounce

import (
	"encoding/json"
)

const (
	ErrorTypeGeneralFailure    string = "general_failure"
	ErrorTypeAuthFailure       string = "auth_failure"
	ErrorTypeBadReferrer       string = "bad_referrer"
	ErrorTypeThrottleTriggered string = "throttle_triggered"
)

type Error struct {
	Type    string `json:"status"`
	Message string `json:"message"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
/*
Package neverbounce creates native Golang mappings to use NeverBounce's email verification API.
Our verification API allows you to create Custom Integrations to add email verification to any part of your software.
We offer solutions for verifying individual emails as well as lists containing hundreds or even millions of emails.

For our full API documentation see: https://developers.neverbounce.com/v4.0/

Basic usage:
	import "github.com/neverbounce/neverbounceapi-go"
	client := neverbounce.New("api_key")

	accountInfo, err := client.Account.Info()
	if err != nil {
		// Attempt to cast the error into a neverbounce.Error to
		// handle-able error objects
		if nbError, ok := err.(*neverbounce.Error); ok {
			// Check Error types
			if nbError.Type == neverbounce.ErrorTypeAuthFailure {
				// The API credentials used are bad, have you reset them recently?
			} else if (nbError.Type == neverbounce.ErrorTypeBadReferrer) {
				// The script is being used from an unauthorized source, you may need to
				// adjust your app's settings to allow it to be used from here
			} else if (nbError.Type == neverbounce.ErrorTypeThrottleTriggered) {
				// Too many requests in a short amount of time, try again shortly or adjust
				// your rate limit settings for this application in the dashboard
			} else {
				// A non recoverable API error occurred check the message for details
			}
		} else {
			// Handle non NeverBounce errors
		}
	}
	fmt.Println(accountInfo)

Additional examples can be found in the examples directory
*/
package neverbounce

import (
	"encoding/json"
)

const (
	// ErrorTypeGeneralFailure is a generic error coming from the API
	ErrorTypeGeneralFailure string = "general_failure"

	// ErrorTypeAuthFailure indicates an issue with the API credentials supplied
	ErrorTypeAuthFailure string = "auth_failure"

	// ErrorTypeBadReferrer indicates that the API is being used from an host that hasn't been authorized
	ErrorTypeBadReferrer string = "bad_referrer"

	// ErrorTypeThrottleTriggered indicates that too many requests have been made in a short amount of time
	ErrorTypeThrottleTriggered string = "throttle_triggered"
)

// Error is the structure of for an NeverBounce API error
type Error struct {
	Type    string `json:"status"`
	Message string `json:"message"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

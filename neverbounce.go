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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// NeverBounce is the root struct of the wrapper.
// This is used to access the specific bindings.
type NeverBounce struct {
	Account *Account
	Single  *Single
	Jobs    *Jobs
	POE     *POE
}

// Version is the current version of the wrapper
const Version = "4.1.0"

// DefaultBaseURL is the default host to make the API requests on
const DefaultBaseURL = "https://api.neverbounce.com"

// New creates a new instance of *NeverBounce. Accepts the api key to use for authentication.
func New(apiKey string) *NeverBounce {
	apiBaseURL := DefaultBaseURL + "/v4.2/"
	r := &NeverBounce{
		Account: &Account{
			apiBaseURL: apiBaseURL,
			apiKey:     apiKey,
		},
		Single: &Single{
			apiBaseURL: apiBaseURL,
			apiKey:     apiKey,
		},
		Jobs: &Jobs{
			apiBaseURL: apiBaseURL,
			apiKey:     apiKey,
		},
		POE: &POE{
			apiBaseURL: apiBaseURL,
			apiKey:     apiKey,
		},
	}

	return r
}

// SetBaseURL will set the url used to make the requests (overrides the DefaultBaseURL constant).
// This method is primarily for internal testing and debugging purposes,
// under normal circumstances it will not be used
func (r *NeverBounce) SetBaseURL(url string) {
	r.Account.apiBaseURL = url
	r.Single.apiBaseURL = url
	r.Jobs.apiBaseURL = url
	r.POE.apiBaseURL = url
}

// SetAPIVersion will set API version used to make requests. It overrides DefaultBaseUrl.
func (r *NeverBounce) SetAPIVersion(apiVersion string) {
	apiBaseURL := DefaultBaseURL + "/" + apiVersion + "/"
	r.SetBaseURL(apiBaseURL)
}

// MakeRequest handles the request and parsing of the responses to and from the API
// It will throw and error when a 4xx/5xx HTTP code is encountered or if
// the API returns an api error.
// See: https://developers.neverbounce.com/v4.0/reference#error-handling
func MakeRequest(method string, url string, data interface{}) ([]byte, error) {
	// Marshal struct into JSON
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Make request
	request, _ := http.NewRequest(method, url, bytes.NewReader(body))
	request.Header.Add("User-Agent", "NeverBounceApi-Go/"+Version)
	request.Header.Add("Content-Type", "application/json")

	// Do request
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// handle 5xx HTTP codes
	if res.StatusCode >= 400 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)

		if res.StatusCode >= 500 {
			return nil, &Error{
				Type: "auth_failure",
				Message: "We were unable to complete your request. " +
					"The following information was supplied: " + buf.String() +
					"\n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])",
			}
		}

		// handle 4xx HTTP codes
		if res.StatusCode >= 400 {
			return nil, &Error{
				Type: "general_failure",
				Message: "We were unable to complete your request. " +
					"The following information was supplied: " + buf.String() +
					"\n\n(Request error [status " + strconv.Itoa(res.StatusCode) + "])",
			}
		}
	}
	// Read body from request
	body, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}

	if strings.Contains(url, "v4/jobs/download") == false && res.Header.Get("Content-Type") != "application/json" {
		return nil, &Error{
			Type: "general_failure",
			Message: "The API responded with a datatype of \"" + res.Header.Get("Content-Type") +
				"\", but \"application/json\" was expected." +
				"\n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])",
		}
	}

	// Handle JSON repsonses
	if res.Header.Get("Content-Type") == "application/json" {

		// check error response
		var nbError Error

		// Unmarshal into error
		err := json.Unmarshal(body, &nbError)
		if err != nil {
			buf := new(bytes.Buffer)
			buf.ReadFrom(res.Body)
			return nil, &Error{
				Type: "general_failure",
				Message: "We were unable to parse the API response. " +
					"The following information was supplied: " + buf.String() +
					"\n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])",
			}
		}

		if nbError.Type != "success" {
			if nbError.Type == ErrorTypeAuthFailure {
				return nil, &Error{
					Type: "auth_failure",
					Message: "We were unable to authenticate your request. " +
						"The following information was supplied: " + nbError.Message +
						"\n\n(" + nbError.Type + ")",
				}
			}

			return nil, &Error{
				Type: nbError.Type,
				Message: "We were unable to complete your request. " +
					"The following information was supplied: " + nbError.Message +
					"\n\n(" + nbError.Type + ")",
			}
		}

		// Return json string
		return body, nil
	}

	// Return plain text responses
	return body, nil
}

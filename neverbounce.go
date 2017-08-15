/*
Package neverbounce creates native Golang mappings to use NeverBounce's email verification API.
Our verification API allows you to create Custom Integrations to add email verification to any part of your software.
We offer solutions for verifying individual emails as well as lists containing hundreds or even millions of emails.

For our full API documentation see: https://developers.neverbounce.com/v4.0/

Basic usage:
	import "github.com/neverbounce/neverbounceapi-go"
	client, err := neverbounce.New("api_key")
	if err != nil {
		panic(err)
	}

	accountInfo, err := client.Account.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(accountInfo)

Additional examples can be found in the examples directory
*/
package neverbounce

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/NeverBounce/NeverBounceApi-Go/models"
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
const Version = "4.0.0"

// DefaultBaseURL is the default host to make the API requests on
const DefaultBaseURL = "https://api.neverbounce.com/v4/"

// New creates a new instance of *NeverBounce. Accepts the api key to use for authentication.
func New(apiKey string) (*NeverBounce) {
	r := &NeverBounce{
		Account: &Account{
			apiBaseURL: DefaultBaseURL,
			apiKey:     apiKey,
		},
		Single: &Single{
			apiBaseURL: DefaultBaseURL,
			apiKey:     apiKey,
		},
		Jobs: &Jobs{
			apiBaseURL: DefaultBaseURL,
			apiKey:     apiKey,
		},
		POE: &POE{
			apiBaseURL: DefaultBaseURL,
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
			return nil, errors.New("We were unable to complete your request. " +
				"The following information was supplied: " + buf.String() +
				"\n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])")
		}

		// handle 4xx HTTP codes
		if res.StatusCode >= 400 {
			return nil, errors.New("We were unable to complete your request. " +
				"The following information was supplied: " + buf.String() +
				"\n\n(Request error [status " + strconv.Itoa(res.StatusCode) + "])")
		}
	}
	// Read body from request
	body, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}

	// Handle JSON repsonses
	if res.Header.Get("Content-Type") == "application/json" {

		// check error response
		var apiError nbModels.APIErrorModel

		// Unmarshal into error
		var err = json.Unmarshal(body, &apiError)
		if err != nil {
			return nil, err
		}

		if apiError.Status != "success" {
			return nil, errors.New("We were unable to complete your request. " +
				"The following information was supplied: " + apiError.Message +
				"\n\n(" + apiError.Status + ")")
		}

		// Return json string
		return body, nil
	}

	// Return plain text responses
	return body, nil
}

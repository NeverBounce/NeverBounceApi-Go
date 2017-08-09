/*
Package neverbounce creates Golang friendly mappings to use NeverBounce's email verification API.

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
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

// NeverBounce : Our verification API allows you to create Custom Integrations to add email verification to any part of your software.
// We offer solutions for verifying individual emails as well as lists containing hundreds or even millions of emails.
type NeverBounce struct {
	Account     *Account
	Single     *Single
	Jobs       *Jobs
}

const DEFAULT_BASE_URL = "https://api.neverbounce.com/v4/"

// New : Create a new instance of *NeverBounce
// @Param
// apiKey: API authentication key
func New(apiKey string) (*NeverBounce, error) {
	r := &NeverBounce{
		Account: &Account{
			apiBaseURL: DEFAULT_BASE_URL,
			apiKey: apiKey,
		},
		Single: &Single{
			apiBaseURL: DEFAULT_BASE_URL,
			apiKey: apiKey,
		},
		Jobs: &Jobs{
			apiBaseURL: DEFAULT_BASE_URL,
			apiKey:     apiKey,
		},
	}

	return r, nil
}

func (r *NeverBounce) SetBaseURL(url string) {
	r.Account.apiBaseURL = url
	r.Single.apiBaseURL = url
	r.Jobs.apiBaseURL = url
}

func makeRequest(method string, url string, data interface{}) ([]byte, error) {
	// Marshal struct into JSON
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Make request
	request, _ := http.NewRequest(method, url, bytes.NewReader(body))
	request.Header.Add("Content-Type", "application/json")

	// Do request
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// handle 5xx HTTP codes
	if res.StatusCode >= 500 {
		return nil, errors.New("We were unable to complete your request. " +
			"The following information was supplied: " + res.Status +
			"\n\n(Request error [status " + strconv.Itoa(res.StatusCode) + "])")
	}

	// handle 4xx HTTP codes
	if res.StatusCode >= 400 {
		return nil, errors.New("We were unable to complete your request. " +
			"The following information was supplied: " + res.Status +
			"\n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])")
	}

	// Read body from request
	body, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}

	// Handle JSON repsonses
	if res.Header.Get("Content-Type") == "application/json" {

		// check error response
		var apiError nbModels.ApiErrorModel

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
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
	"encoding/json"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

// POE endpoints allow you to confirm frontend verifications performed
// by the Javascript Widget
type POE struct {
	apiBaseURL string
	apiKey     string
}

// Confirm verifies that the result provided during frontend verification (e.g. Javascript Widget) has
// not been tampered with.
// It requires you to pass the email, result, transaction_id, and confirmation_token supplied by the verification.
func (r *POE) Confirm(model *nbModels.POEConfirmRequestModel) (*nbModels.POEConfirmResponseModel, error) {
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "poe/confirm"
	body, err := MakeRequest("GET", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.POEConfirmResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

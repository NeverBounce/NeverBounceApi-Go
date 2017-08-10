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

// Single endpoints allow you to integrate our email verification into your existing
// applications at the point of entry and onboarding processes
type Single struct {
	apiBaseURL string
	apiKey     string
}

// Check verifies the email provided and returns the verification result.
// In addition to this, it can also return a breakdown of the email address' host info
// and your account balance
func (r *Single) Check(model *nbModels.SingleCheckRequestModel) (*nbModels.SingleCheckResponseModel, error) {
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "single/check"
	body, err := MakeRequest("GET", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.SingleCheckResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

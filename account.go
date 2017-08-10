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

// Account contains bindings for account related API endpoints.
type Account struct {
	apiBaseURL string
	apiKey     string
}

// Info returns the account's current credit balance as well as job counts
// indicating the number of jobs currently in the account.
func (r *Account) Info() (*nbModels.AccountInfoResponseModel, error) {
	// call info API
	url := r.apiBaseURL + "account/info"
	body, err := MakeRequest("GET", url, &nbModels.GenericRequestModel{APIKey: r.apiKey})
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.AccountInfoResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

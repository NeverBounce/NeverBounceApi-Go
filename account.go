// Package neverBounce wrap NeverBounce restful APIs
package neverbounce

import (
	"encoding/json"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

// Single : Single functionality holder
type Account struct {
	apiBaseURL string
	apiKey     string
}

// Info : Account endpoints allow to programmatically check your account's balance and
// how many jobs are currently running on your account.
func (r *Account) Info() (*nbModels.AccountInfoResponseModel, error) {
	// call info API
	url := r.apiBaseURL + "account/info"
	body, err := makeRequest("GET", url, &nbModels.GenericRequestModel{ApiKey: r.apiKey})
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

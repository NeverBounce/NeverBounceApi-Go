package neverbounce

import (
	"encoding/json"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

// Contains bindings for Account API endpoints.
type Account struct {
	apiBaseURL string
	apiKey     string
}

// Account info returns the account's current credit balance as well as job counts
// indicating the number of jobs currently in the account.
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

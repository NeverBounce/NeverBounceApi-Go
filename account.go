// Package neverBounce wrap NeverBounce restful APIs
package neverbounce

import (
	"encoding/json"
	"errors"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_dto"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_error"
)

// Info : Account endpoints allow to programmatically check your account's balance and
// how many jobs are currently running on your account.
func (r *NeverBounce) Info() (*nbModels.AccountInfo, error) {
	// call info API
	url := r.apiBaseURL + "account/info?key=" + r.APIKey

	body, err := callAPI(url)
	if err != nil {
		return nil, err
	}

	// check error response
	var apiError nbErrors.ApiError

	err = json.Unmarshal(body, &apiError)
	if err != nil {
		return nil, err
	}
	if apiError.Status == "auth_failure" {
		return nil, errors.New(apiError.Message)
	}

	// check success response
	var info nbModels.AccountInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

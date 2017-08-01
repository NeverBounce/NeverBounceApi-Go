// Package neverBounce wrap NeverBounce restful APIs
package neverBounce

import (
	"encoding/json"
	"errors"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_dto"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_error"
)

// Info : Account endpoints allow to programmatically check your account's balance and
// how many jobs are currently running on your account.
func (r *NeverBounce) Info() (*nbDto.AccountInfo, error) {
	// call info API
	url := r.apiBaseURL + "account/info?key=" + r.APIKey

	body, err := callAPI(url)
	if err != nil {
		return nil, err
	}

	// check error response
	var authError nbError.AuthError

	err = json.Unmarshal(body, &authError)
	if err != nil {
		return nil, err
	}
	if authError.Status == "auth_failure" {
		return nil, errors.New(authError.Message)
	}

	// check success response
	var info nbDto.AccountInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

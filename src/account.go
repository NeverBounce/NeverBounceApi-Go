package neverBounce

import (
	"encoding/json"
	"errors"
	"github.com/enkhalifapro/neverBounce/src/nb_error"
	"github.com/enkhalifapro/neverBounce/src/nb_dto"
)

// Account endpoints allow to programmatically check your account's balance and
// how many jobs are currently running on your account.
func (r *NeverBounce) Info() (error, *nbDto.AccountInfo) {
	// call info API
	url := r.apiBaseUrl + "account/info?key=" + r.ApiKey

	err, body := callApi(url)
	if err != nil {
		return err, nil
	}

	// check error response
	var authError nbError.AuthError

	err = json.Unmarshal(body, &authError)
	if err != nil {
		return err, nil
	}
	if authError.Status == "auth_failure" {
		return errors.New(authError.Message), nil
	}

	// check success response
	var info nbDto.AccountInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return err, nil
	}
	return nil, &info
}

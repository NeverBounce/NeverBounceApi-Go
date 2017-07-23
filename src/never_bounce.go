package neverBounce

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"github.com/enkhalifapro/neverBounce/src/nb_error"
	"github.com/enkhalifapro/neverBounce/src/nb_dto"
)

type NeverBounce struct {
	ApiKey     string
	apiBaseUrl string
}

func New(apiKey string) (error, *NeverBounce) {
	r := &NeverBounce{}
	r.apiBaseUrl = "https://api.neverbounce.com/v4/"
	r.ApiKey = apiKey
	err, _ := r.Info()
	if err != nil {
		return err, nil
	}
	return nil, r
}

func callApi(url string) (error, []byte) {
	res, err := http.Get(url)
	if err != nil {
		return err, nil
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err, nil
	}
	return nil, body
}

// Account endpoints allow to programmatically check your account's balance and
// how many jobs are currently running on your account.
func (r *NeverBounce) Info() (error, *nbDto.Info) {
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
	var info nbDto.Info

	err = json.Unmarshal(body, &info)
	if err != nil {
		return err, nil
	}
	return nil, &info
}

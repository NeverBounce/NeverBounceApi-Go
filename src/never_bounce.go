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
	ApiKey string
}

func New(apiKey string) (error, *NeverBounce) {
	r := &NeverBounce{}
	r.ApiKey = apiKey
	return nil, r
}

func (r *NeverBounce) Info() (error, *nbDto.Info) {
	url := "https://api.neverbounce.com/v4/account/info?key=" + r.ApiKey
	res, err := http.Get(url)
	if err != nil {
		return err, nil
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err, nil
	}
	var authError nbError.AuthError

	err = json.Unmarshal(body, &authError)
	if err != nil {
		return err, nil
	}
	if authError.Status == "auth_failure" {
		return errors.New(authError.Message), nil
	}

	var info nbDto.Info

	err = json.Unmarshal(body, &info)
	if err != nil {
		return err, nil
	}
	return nil, &info
}

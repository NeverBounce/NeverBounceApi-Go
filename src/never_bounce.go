package neverBounce

import (
	"net/http"
	"io/ioutil"
)

type NeverBounce struct {
	apiBaseUrl string
	ApiKey     string
	Single     *Single
}

func New(apiKey string) (error, *NeverBounce) {
	baseUrl := "https://api.neverbounce.com/v4/"
	r := &NeverBounce{
		apiBaseUrl: baseUrl,
		ApiKey:     apiKey,
		Single: &Single{apiBaseUrl: baseUrl,
			apiKey:                 apiKey}}
	_, err := r.Info()
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

// Package neverBounce wrap NeverBounce restful APIs
package neverBounce

import (
	"net/http"
	"io/ioutil"
)

// NeverBounce : Our verification API allows you to create Custom Integrations to add email verification to any part of your software.
// We offer solutions for verifying individual emails as well as lists containing hundreds or even millions of emails.
type NeverBounce struct {
	apiBaseURL string
	APIKey     string
	Single     *Single
}

// New : Create a new instance of *NeverBounce
// @Param
// apiKey: API authentication key
func New(apiKey string) (*NeverBounce, error) {
	baseURL := "https://api.neverbounce.com/v4/"
	r := &NeverBounce{
		apiBaseURL: baseURL,
		APIKey:     apiKey,
		Single: &Single{apiBaseURL: baseURL,
			apiKey:                 apiKey}}
	_, err := r.Info()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func callAPI(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

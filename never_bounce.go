// Package neverBounce wrap NeverBounce restful APIs
package neverBounce

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"os"
	"io"
	"encoding/json"
	"errors"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_error"
	"strconv"
)

// NeverBounce : Our verification API allows you to create Custom Integrations to add email verification to any part of your software.
// We offer solutions for verifying individual emails as well as lists containing hundreds or even millions of emails.
type NeverBounce struct {
	apiBaseURL string
	APIKey     string
	Single     *Single
	Jobs       *Jobs
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
			apiKey:                 apiKey},
		Jobs: &Jobs{
			apiBaseURL: baseURL,
			apiKey:     apiKey}}
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

	// handle 4xx HTTP codes
	if res.StatusCode >= 500 {
		return nil, errors.New("We were unable to complete your request. The following information was supplied \n\n(Request error [status " + strconv.Itoa(res.StatusCode) + "])")
	}
	// handle 5xx HTTP codes
	if res.StatusCode >= 400 && res.StatusCode < 500 {
		return nil, errors.New("We were unable to complete your request. The following information was supplied \n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func postAPI(url string, postedBody *bytes.Buffer) ([]byte, error) {
	res, err := http.Post(url, "application/json", postedBody)
	if err != nil {
		return nil, err
	}

	// handle 4xx HTTP codes
	if res.StatusCode >= 500 {
		return nil, errors.New("We were unable to complete your request. The following information was supplied \n\n(Request error [status " + strconv.Itoa(res.StatusCode) + "])")
	}
	// handle 5xx HTTP codes
	if res.StatusCode >= 400 && res.StatusCode < 500 {
		return nil, errors.New("We were unable to complete your request. The following information was supplied \n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func downloadFile(filepath string, url string) (err error) {
	// Get the data
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	// handle 4xx HTTP codes
	if res.StatusCode >= 500 {
		return errors.New("We were unable to complete your request. The following information was supplied \n\n(Request error [status " + strconv.Itoa(res.StatusCode) + "])")
	}
	// handle 5xx HTTP codes
	if res.StatusCode >= 400 && res.StatusCode < 500 {
		return errors.New("We were unable to complete your request. The following information was supplied \n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// check error response
	var authError nbError.AuthError
	err = json.Unmarshal(body, &authError)
	if err != nil {
		return err
	}
	if authError.Status != "success" {
		return errors.New(authError.Message)
	}

	// Writer the body to file
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}

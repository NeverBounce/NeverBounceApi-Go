package neverBounce

import (
	"encoding/json"
	"github.com/NeverBounce/NeverBounceApi-Go/src/nb_dto"
)

// Single : Single functionality holder
type Single struct {
	apiBaseURL string
	apiKey     string
}

// Check : verification allows you verify individual emails and gather additional
// information pertaining to the email.
// @Params
// email: The email to verify
// includeAddressInfo: Include additional address info in response
// includeCreditInfo: Include account credit info in response
// max_execution_time: The maximum time in seconds we should try to verify the address
func (r *Single) Check(email string, includeAddressInfo bool, includeCreditInfo bool, maxExecutionTime string) (*nbDto.SingleCheckInfo, error) {
	// call info API
	url := r.apiBaseURL + "single/check?key=" + r.apiKey + "&email=" + email

	// include address info
	if includeAddressInfo == true {
		url += "&address_info=1"
	}

	// include credit info
	if includeAddressInfo == true {
		url += "&credits_info=1"
	}

	// include maxExecutionTime
	if maxExecutionTime != "" {
		url += "&max_execution_time=" + maxExecutionTime
	}

	body, err := callAPI(url)
	if err != nil {
		return nil, err
	}

	// extract result info
	var info nbDto.SingleCheckInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Package neverBounce wrap NeverBounce restful APIs
package neverbounce

import (
	"encoding/json"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
	"github.com/NeverBounce/NeverBounceApi-Go/errors"
	"errors"
)

// Single : Single functionality holder
type Single struct {
	apiBaseURL string
	apiKey     string
}

// Check : verification allows you verify individual emails and gather additional
// information pertaining to the email.
func (r *Single) Check(email string, includeAddressInfo bool, includeCreditInfo bool, maxExecutionTime string) (*nbModels.SingleCheckModel, error) {
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

	// check error response
	var apiError nbErrors.ApiError
	err = json.Unmarshal(body, &apiError)
	if err != nil {
		return nil, err
	}
	if apiError.Status != "success" {
		return nil, errors.New(apiError.Message)
	}

	// extract result info
	var info nbModels.SingleCheckModel

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

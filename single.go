// Package neverBounce wrap NeverBounce restful APIs
package neverbounce

import (
	"encoding/json"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

// Single : Single functionality holder
type Single struct {
	apiBaseURL string
	apiKey     string
}

// Check : verification allows you verify individual emails and gather additional
// information pertaining to the email.
func (r *Single) Check(model *nbModels.SingleCheckRequestModel) (*nbModels.SingleCheckResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "single/check"
	body, err := makeRequest("GET", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.SingleCheckResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

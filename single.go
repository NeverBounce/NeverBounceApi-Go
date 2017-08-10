// Package neverBounce wrap NeverBounce restful APIs
package neverbounce

import (
	"encoding/json"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

// Single endpoints allow you to integrate our email verification into your existing
// applications at the point of entry and onboarding processes
type Single struct {
	apiBaseURL string
	apiKey     string
}

// Single check verifies the email provided and returns the verification result.
// In addition to this, it can also return a breakdown of the email address' host info
// and your account balance
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

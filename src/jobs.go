// Package neverBounce wrap NeverBounce restful APIs
package neverBounce

import (
	"encoding/json"
	"strconv"
	"github.com/NeverBounce/NeverBounceApi-Go/src/nb_dto"
	"bytes"
	"github.com/NeverBounce/NeverBounceApi-Go/src/nb_error"
	"errors"
)

// Jobs : The bulk endpoint provides high-speed​ validation on a list of email addresses.
// You can use the status endpoint to retrieve real-time statistics about a bulk job in progress.
// Once the job has finished, the results can be retrieved with our download endpoint.
type Jobs struct {
	apiBaseURL string
	apiKey     string
}

// Search : filter and find from the saved validation jobs
func (r *Jobs) Search(jobID int, fileName string, completed bool, processing bool, indexing bool, failed bool, manualReview bool, unpurchased bool, page int, itemsPerPage int) (*nbDto.SearchInfo, error) {
	// call API
	url := r.apiBaseURL + "jobs/search?key=" + r.apiKey

	// add jobId param
	if jobID > 0 {
		url += "&job_id=" + strconv.Itoa(jobID)
	}

	// add completed param
	if completed != false {
		url += "&completed=1"
	}

	// add processing param
	if processing != false {
		url += "&processing=1"
	}

	// add indexing param
	if indexing != false {
		url += "&indexing=1"
	}

	// add indexing param
	if failed != false {
		url += "&failed=1"
	}

	// add manual_review param
	if manualReview != false {
		url += "&manual_review=1"
	}

	// add unpurchased param
	if unpurchased != false {
		url += "&unpurchased=1"
	}

	// add page param
	if page > 0 {
		url += "&page=" + strconv.Itoa(jobID)
	}

	// add page param
	if itemsPerPage > 0 {
		url += "&items_per_page=" + strconv.Itoa(jobID)
	}

	body, err := callAPI(url)
	if err != nil {
		return nil, err
	}

	// extract result info
	var info nbDto.SearchInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Create : add a new validation job
// @Param
// inputLocation: The type of input being supplied. Accepted values are "remote_url" and "supplied".
// input: The input to be verified
// autoParse: Should be begin parsing the job immediately?
// autoRun: Should we run the job immediately after being parsed?
// runSample: Should this job be run as a sample?
// fileName: This will be what's displayed in the dashboard when viewing this job
func (r *Jobs) Create(createSearch *nbDto.CreateSearch) (*nbDto.CreateSearchInfo, error) {
	// call API
	url := r.apiBaseURL + "jobs/create"
	createSearch.APIKEY = r.apiKey
	postedBody, err := json.Marshal(createSearch)
	if err != nil {
		return nil, err
	}
	body, err := postAPI(url, bytes.NewBuffer(postedBody))
	if err != nil {
		return nil, err
	}

	// check error response
	var authError nbError.AuthError
	err = json.Unmarshal(body, &authError)
	if err != nil {
		return nil, err
	}
	if authError.Status != "success" {
		return nil, errors.New(authError.Message)
	}

	// extract result info
	var info nbDto.CreateSearchInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

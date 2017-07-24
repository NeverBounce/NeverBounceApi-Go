package neverBounce

import (
	"encoding/json"
	"github.com/enkhalifapro/neverBounce/src/nb_dto"
	"strconv"
)

type Jobs struct {
	apiBaseUrl string
	apiKey     string
}

// Single verification allows you verify individual emails and gather additional
// information pertaining to the email.
// @Params
// email: The email to verify
// includeAddressInfo: Include additional address info in response
// includeCreditInfo: Include account credit info in response
// max_execution_time: The maximum time in seconds we should try to verify the address

func (r *Jobs) Search(jobId int, fileName string, completed bool, processing bool, indexing bool, failed bool, manualReview bool, unpurchased bool, page int, itemsPerPage int) (error, *nbDto.SingleCheckInfo) {
	// call info API
	url := r.apiBaseUrl + "jobs/search?key=" + r.apiKey

	// add jobId param
	if jobId > 0 {
		url += "&job_id=" + strconv.Itoa(jobId)
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
		url += "&page=" + strconv.Itoa(jobId)
	}

	// add page param
	if itemsPerPage > 0 {
		url += "&items_per_page=" + strconv.Itoa(jobId)
	}

	err, body := callApi(url)
	if err != nil {
		return err, nil
	}

	// extract result info
	var info nbDto.SingleCheckInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return err, nil
	}
	return nil, &info
}

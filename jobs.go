// Package neverBounce wrap NeverBounce restful APIs
package neverBounce

import (
	"encoding/json"
	"strconv"
	"bytes"
	"errors"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_error"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_dto"
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

// Parse : allows you to parse a job created with auto_parse disabled.
// You cannot reparse a list once it's been parsed.
func (r *Jobs) Parse(jobID int, autoStart bool) (*nbDto.ParseInfo, error) {
	// call API
	url := r.apiBaseURL + "jobs/parse"
	values := map[string]interface{}{}
	values["key"] = r.apiKey
	values["job_id"] = jobID
	values["auto_start"] = autoStart
	postedBody, err := json.Marshal(values)
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
	var info nbDto.ParseInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Start : allows you to start a job created or parsed with auto_start disabled.
// Once the list has been started the credits will be deducted and the process cannot be stopped or restarted
func (r *Jobs) Start(jobID int, runSample bool) (*nbDto.StartInfo, error) {
	// call API
	url := r.apiBaseURL + "jobs/start"
	values := map[string]interface{}{}
	values["key"] = r.apiKey
	values["job_id"] = jobID
	values["run_sample"] = runSample
	postedBody, err := json.Marshal(values)
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
	var info nbDto.StartInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Status : indicate what stage the job is currently in.
// This will be the primary property you'll want to check to determine what can be done with the job.
func (r *Jobs) Status(jobID int) (*nbDto.JobStatusInfo, error) {
	// call API
	url := r.apiBaseURL + "jobs/status?key=" + r.apiKey + "&job_id=" + strconv.Itoa(jobID)
	body, err := callAPI(url)
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
	var info nbDto.JobStatusInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Results : Get job result by job ID.
func (r *Jobs) Results(jobID int, page int, itemPerPage int) (*nbDto.ResultInfo, error) {
	// call API
	url := r.apiBaseURL + "jobs/results?key=" + r.apiKey + "&job_id=" + strconv.Itoa(jobID)

	// add page param
	if page > 0 {
		url += "&page=" + strconv.Itoa(page)
	}

	// add itemPerPage param
	if page > 0 {
		url += "&items_per_page=" + strconv.Itoa(itemPerPage)
	}

	body, err := callAPI(url)
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
	var info nbDto.ResultInfo

	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Download : Download a file containing the job data as a CSV file.
func (r *Jobs) Download(jobID int, filePath string) (error) {
	// call API
	url := r.apiBaseURL + "jobs/download?key=" + r.apiKey + "&job_id=" + strconv.Itoa(jobID)
	err := downloadFile(filePath, url)
	return err
}

// DownloadWithOptions : Download a file containing the job data as a CSV file with using filters.
func (r *Jobs) DownloadWithOptions(jobID int, filePath string, valids int, inValids int, catchalls int, unknowns int, disposables int, includeDuplicates int, emailStatus int) (error) {
	// call API
	url := r.apiBaseURL + "jobs/download?key=" + r.apiKey + "&job_id=" + strconv.Itoa(jobID)
	url += "&valids=" + strconv.Itoa(valids)
	url += "&invalids=" + strconv.Itoa(inValids)
	url += "&catchalls=" + strconv.Itoa(catchalls)
	url += "&unknowns=" + strconv.Itoa(unknowns)
	url += "&disposables=" + strconv.Itoa(disposables)
	url += "&include_duplicates=" + strconv.Itoa(includeDuplicates)
	url += "&email_status=" + strconv.Itoa(emailStatus)
	err := downloadFile(filePath, url)
	return err
}

// Delete : delete job by ID
func (r *Jobs) Delete(jobID int) (error) {
	// call API
	url := r.apiBaseURL + "jobs/delete"
	values := map[string]interface{}{}
	values["key"] = r.apiKey
	values["job_id"] = jobID
	postedBody, err := json.Marshal(values)
	if err != nil {
		return err
	}
	body, err := postAPI(url, bytes.NewBuffer(postedBody))
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

	return nil
}

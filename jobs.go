// Package neverBounce wrap NeverBounce restful APIs
package neverbounce

import (
	"encoding/json"
	"bytes"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
	"io"
	"os"
)

// Jobs : The bulk endpoint provides high-speed​ validation on a list of email addresses.
// You can use the status endpoint to retrieve real-time statistics about a bulk job in progress.
// Once the job has finished, the results can be retrieved with our download endpoint.
type Jobs struct {
	apiBaseURL string
	apiKey     string
}

// Search : filter and find from the saved validation jobs
func (r *Jobs) Search(model *nbModels.JobsSearchRequestModel) (*nbModels.JobsSearchResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/search"
	body, err := makeRequest("GET", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.JobsSearchResponseModel
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
func (r *Jobs) CreateFromSuppliedData(model *nbModels.JobsCreateSuppliedDataRequestModel) (*nbModels.JobsCreateResponseModel, error) {
	model.ApiKey = r.apiKey
	model.InputLocation = "supplied"

	// call info API
	url := r.apiBaseURL + "jobs/create"
	body, err := makeRequest("POST", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.JobsCreateResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *Jobs) CreateFromRemoteUrl(model *nbModels.JobsCreateRemoteUrlRequestModel) (*nbModels.JobsCreateResponseModel, error) {
	model.ApiKey = r.apiKey
	model.InputLocation = "remote_url"

	// call info API
	url := r.apiBaseURL + "jobs/create"
	body, err := makeRequest("POST", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.JobsCreateResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Parse : allows you to parse a job created with auto_parse disabled.
// You cannot reparse a list once it's been parsed.
func (r *Jobs) Parse(model *nbModels.JobsParseRequestModel) (*nbModels.JobsParseResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/parse"
	body, err := makeRequest("POST", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.JobsParseResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Start : allows you to start a job created or parsed with auto_start disabled.
// Once the list has been started the credits will be deducted and the process cannot be stopped or restarted
func (r *Jobs) Start(model *nbModels.JobsStartRequestModel) (*nbModels.JobsStartResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/start"
	body, err := makeRequest("POST", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.JobsStartResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Status : indicate what stage the job is currently in.
// This will be the primary property you'll want to check to determine what can be done with the job.
func (r *Jobs) Status(model *nbModels.JobsStatusRequestModel) (*nbModels.JobsStatusResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/status"
	body, err := makeRequest("POST", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.JobsStatusResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Results : Get job result by job ID.
func (r *Jobs) Results(model *nbModels.JobsResultsRequestModel) (*nbModels.JobsResultsResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/results"
	body, err := makeRequest("POST", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.JobsResultsResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Download : Download a file containing the job data as a CSV file.
func (r *Jobs) Download(model *nbModels.JobsDownloadRequestModel, filepath string) (error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/download"
	body, err := makeRequest("POST", url, model)
	if err != nil {
		return err
	}

	// Writer the body to file
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, bytes.NewReader(body))
	return err
}

// Delete : delete job by ID
func (r *Jobs) Delete(model *nbModels.JobsDeleteRequestModel) (*nbModels.JobsDeleteResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/delete"
	body, err := makeRequest("POST", url, model)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var info nbModels.JobsDeleteResponseModel
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}


package neverbounce

import (
	"encoding/json"
	"bytes"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
	"io"
	"os"
)

// The jobs endpoints provides high-speed​ validation on a list of email addresses.
// You can use the status endpoint to retrieve real-time statistics about a bulk job in progress.
// Once the job has finished, the results can be retrieved with our download endpoint.
type Jobs struct {
	apiBaseURL string
	apiKey     string
}

// Search the jobs you've previously submitted to your account.
// It will return jobs in batches according to the pagination options you've supplied.
// The returned jobs will include the same information available from the Status method
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

// Creates a new job from data you supply directly in the request.
// Supplied data will need to be given as a map, see the examples in the nbModel package.
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

// Creates a new job from a comma separated value (CSV) file hosted on a remote URL.
// The URL supplied can be any commonly available protocal; e.g: HTTP, HTTPS, FTP, SFTP.
// Basic auth is supported by including the credentials in the URI string; e.g: http://name:passwd@example.com/full/path/to/file.csv
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

// If you create a job with AutoParse set to false (defaults to false) you can parse job using this endpoint.
// Once parsed, a job cannot be reparsed.
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

// If you create a job or parse a job with AutoStart set to false (defaults to false) you can start the job with this method.
// Once the list has been started the credits will be deducted and the process cannot be stopped or restarted.
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

// Status will return information pertaining to the Jobs state. It will include the jobs current status as well as the verification stats.
// This will be the primary property you'll want to check to determine what can be done with the job.
func (r *Jobs) Status(model *nbModels.JobsStatusRequestModel) (*nbModels.JobsStatusResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/status"
	body, err := makeRequest("GET", url, model)
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

// Results will return the actual verification results.
// This can only be done once the job has reached the completed status.
// The results will be returned in batches according to the pagination options you've supplied.
// Verification info will be formatted the same way Single.Check returns verification info.
func (r *Jobs) Results(model *nbModels.JobsResultsRequestModel) (*nbModels.JobsResultsResponseModel, error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/results"
	body, err := makeRequest("GET", url, model)
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

// Download the results as a CSV to a file.
// This is useful if your uploading the results to a CRM or are use to working with spreadsheets.
func (r *Jobs) Download(model *nbModels.JobsDownloadRequestModel, filepath string) (error) {
	model.ApiKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/download"
	body, err := makeRequest("GET", url, model)
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

// Delete a job. This can only be done when a job is Queued, Waiting, Completed, or Failed.
// A job cannot be deleted while it is being uploaded, parsed, or ran.
// Once deleted the job results cannot be recovered, deletion is permanent.
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


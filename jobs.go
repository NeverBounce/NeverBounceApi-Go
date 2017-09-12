/*
Package neverbounce creates native Golang mappings to use NeverBounce's email verification API.
Our verification API allows you to create Custom Integrations to add email verification to any part of your software.
We offer solutions for verifying individual emails as well as lists containing hundreds or even millions of emails.

For our full API documentation see: https://developers.neverbounce.com/v4.0/

Basic usage:
	import "github.com/neverbounce/neverbounceapi-go"
	client := neverbounce.New("api_key")

	accountInfo, err := client.Account.Info()
	if err != nil {
		// Attempt to cast the error into a neverbounce.Error to
		// handle-able error objects
		if nbError, ok := err.(*neverbounce.Error); ok {
			// Check Error types
			if nbError.Type == neverbounce.ErrorTypeAuthFailure {
				// The API credentials used are bad, have you reset them recently?
			} else if (nbError.Type == neverbounce.ErrorTypeBadReferrer) {
				// The script is being used from an unauthorized source, you may need to
				// adjust your app's settings to allow it to be used from here
			} else if (nbError.Type == neverbounce.ErrorTypeThrottleTriggered) {
				// Too many requests in a short amount of time, try again shortly or adjust
				// your rate limit settings for this application in the dashboard
			} else {
				// A non recoverable API error occurred check the message for details
			}
		} else {
			// Handle non NeverBounce errors
		}
	}
	fmt.Println(accountInfo)

Additional examples can be found in the examples directory
*/
package neverbounce

import (
	"encoding/json"
	"bytes"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
	"io"
	"os"
)

// Jobs endpoints provides high-speed​ validation on a list of email addresses.
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
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/search"
	body, err := MakeRequest("GET", url, model)
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

// CreateFromSuppliedData creates a new job from data you supply directly in the request.
// Supplied data will need to be given as a map, see the examples in the nbModel package.
func (r *Jobs) CreateFromSuppliedData(model *nbModels.JobsCreateSuppliedDataRequestModel) (*nbModels.JobsCreateResponseModel, error) {
	model.APIKey = r.apiKey
	model.InputLocation = "supplied"

	// call info API
	url := r.apiBaseURL + "jobs/create"
	body, err := MakeRequest("POST", url, model)
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

// CreateFromRemoteURL creates a new job from a comma separated value (CSV) file hosted on a remote URL.
// The URL supplied can be any commonly available protocal; e.g: HTTP, HTTPS, FTP, SFTP.
// Basic auth is supported by including the credentials in the URI string; e.g: http://name:passwd@example.com/full/path/to/file.csv
func (r *Jobs) CreateFromRemoteURL(model *nbModels.JobsCreateRemoteURLRequestModel) (*nbModels.JobsCreateResponseModel, error) {
	model.APIKey = r.apiKey
	model.InputLocation = "remote_url"

	// call info API
	url := r.apiBaseURL + "jobs/create"
	body, err := MakeRequest("POST", url, model)
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

// Parse allows you to parse the job data after creation.
// If you create a job with AutoParse set to true (defaults to false) you do not need to use this method.
// Once parsed, a job cannot be reparsed.
func (r *Jobs) Parse(model *nbModels.JobsParseRequestModel) (*nbModels.JobsParseResponseModel, error) {
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/parse"
	body, err := MakeRequest("POST", url, model)
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

// Start allows you to start a job after it has been parsed.
// If you create a job or parse a job with AutoStart set to true (defaults to false) you do not need to use this method.
// Once the list has been started the credits will be deducted and the process cannot be stopped or restarted.
func (r *Jobs) Start(model *nbModels.JobsStartRequestModel) (*nbModels.JobsStartResponseModel, error) {
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/start"
	body, err := MakeRequest("POST", url, model)
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
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/status"
	body, err := MakeRequest("GET", url, model)
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
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/results"
	body, err := MakeRequest("GET", url, model)
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
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/download"
	body, err := MakeRequest("GET", url, model)
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

// Delete will remove the job and it's verification data (if previously verified)
// This can only be done when a job is Queued, Waiting, Completed, or Failed.
// A job cannot be deleted while it is being uploaded, parsed, or ran.
// Once deleted the job results cannot be recovered, deletion is permanent.
func (r *Jobs) Delete(model *nbModels.JobsDeleteRequestModel) (*nbModels.JobsDeleteResponseModel, error) {
	model.APIKey = r.apiKey

	// call info API
	url := r.apiBaseURL + "jobs/delete"
	body, err := MakeRequest("POST", url, model)
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


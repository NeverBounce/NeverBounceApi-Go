// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// JobsStatusRequestModel is the request model for Jobs.Status()
type JobsStatusRequestModel struct {
	GenericRequestModel
	JobID       int `json:"job_id"`
}

// JobsStatusResponseModel is the response model for Jobs.Status()
type JobsStatusResponseModel struct {
	GenericResponseModel
	JobStatusModel
}

// JobStatusModel is the model for the job's information
type JobStatusModel struct {
	JobID           int `json:"id"`
	FileName        string `json:"filename"`
	CreatedAt       string `json:"created_at"`
	StartedAt       string `json:"started_at"`
	FinishedAt      string `json:"finished_at"`
	Totals          JobStatusTotalsModel `json:"total"`
	BounceEstimate  int `json:"bounce_estimate"`
	PercentComplete int `json:"percent_complete"`
}

// JobStatusTotalsModel is the model for the job's stats
type JobStatusTotalsModel struct {
	Records    int `json:"records"`
	Billable   int `json:"billable"`
	Processed  int `json:"processed"`
	Valid      int `json:"valid"`
	Invalid    int `json:"invalid"`
	Catchall   int `json:"catchall"`
	Disposable int `json:"disposable"`
	Unknown    int `json:"unknown"`
	Duplicates int `json:"duplicates"`
	BadSyntax  int `json:"bad_syntax"`
}
package nbModels

type JobsStatusRequestModel struct {
	GenericRequestModel
	JobId       int `json:"job_id"`
}

type JobsStatusResponseModel struct {
	GenericResponseModel
	JobStatusModel
}

// Jobs status model
type JobStatusModel struct {
	JobId           int `json:"id"`
	FileName        string `json:"filename"`
	CreatedAt       string `json:"created_at"`
	StartedAt       string `json:"started_at"`
	FinishedAt      string `json:"finished_at"`
	Totals          JobStatusTotalsModel `json:"total"`
	BounceEstimate  int `json:"bounce_estimate"`
	PercentComplete int `json:"percent_complete"`
}

// Job totals model
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
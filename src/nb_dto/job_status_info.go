// Package nbDto holds API result models.
package nbDto

// JobStatusInfo : Result model of Info API
type JobStatusInfo struct {
	Status          string `json:"status"`
	ID              string `json:"id"`
	FileName        string `json:"filename"`
	Created         string `json:"created"`
	Started         string `json:"started"`
	Finished        string `json:"finished"`
	TotalRecords    int `json:"total_records"`
	TotalBillable   int `json:"total_billable"`
	TotalProcessed  int `json:"total_processed"`
	TotalValid      int `json:"total_valid"`
	TotalInvalid    int `json:"total_invalid"`
	TotalCatchall   int `json:"total_catchall"`
	TotalDisposable int `json:"total_disposable"`
	TotalDuplicates int `json:"total_duplicates"`
	TotalBadSyntax  int `json:"total_bad_syntax"`
	BounceEstimate  int `json:"bounce_estimate"`
	TotalUnknown    int `json:"total_unknown"`
	PercentComplete int `json:"percent_complete"`
	JobStatus       string `json:"job_status"`
	ExecutionTime   int `json:"execution_time"`
}

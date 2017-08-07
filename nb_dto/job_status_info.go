// Package nbDto holds API result models.
package nbDto

// JobStatusTotal : Nested result model of Status API
type JobStatusTotal struct {
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

// JobStatusInfo : Result model of Status API
type JobStatusInfo struct {
	Status          string `json:"status"`
	ID              int `json:"id"`
	FileName        string `json:"filename"`
	Created         string `json:"created_at"`
	Started         string `json:"started_at"`
	Finished        string `json:"finished_at"`
	Total           JobStatusTotal `json:"total"`
	BounceEstimate  int `json:"bounce_estimate"`
	PercentComplete int `json:"percent_complete"`
	JobStatus       string `json:"job_status"`
	ExecutionTime   int `json:"execution_time"`
}

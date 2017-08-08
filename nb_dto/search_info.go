// Package nbDto holds API result models.
package nbModels

// Query : Nested result model of Info API
type Query struct {
	Page         int `json:"page"`
	ItemsPerPage int `json:"items_per_page"`
}

// SearchResult : Result model of Search API
type SearchResult struct {
	ID              int `json:"id"`
	Status          int `json:"items_per_page"`
	FileName        int `json:"filename"`
	Created         string `json:"created"`
	Started         string `json:"started"`
	Finished        int `json:"finished"`
	TotalRecords    int `json:"total_records"`
	TotalBillable   int `json:"total_billable"`
	TotalProcessed  int `json:"total_processed"`
	TotalValid      int `json:"total_valid"`
	TotalInvalid    int `json:"total_invalid"`
	TotalCatchall   int `json:"total_catchall"`
	TotalDisposable int `json:"total_disposable"`
	TotalUnknown    int `json:"total_unknown"`
	TotalDuplicates int `json:"total_duplicates"`
	TotalBadSyntax  int `json:"total_bad_syntax"`
	BounceEstimate  int `json:"bounce_estimate"`
	PercentComplete int `json:"percent_complete"`
}

// SearchInfo : Nested result model of Info API
type SearchInfo struct {
	Status        string `json:"status"`
	TotalResults  int `json:"total_results"`
	TotalPages    int `json:"total_pages"`
	Query         Query `json:"query"`
	Results       []SearchResult `json:"results"`
	ExecutionTime int `json:"execution_time"`
}

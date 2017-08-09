package nbModels

// Job Search request model
type JobsSearchRequestModel struct {
	GenericRequestModel
	ItemsPerPage int `json:"items_per_page,omitempty"`
	Page         int `json:"page,omitempty"`
	JobId        int `json:"job_id,omitempty"`
	JobFilename  string `json:"filename,omitempty"`
	JobStatus    string `json:"job_status,omitempty"`
}

// Job Search response
type JobsSearchResponseModel struct {
	GenericResponseModel
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Query        JobsSearchQueryModel `json:"query"`
	Results      []JobStatusModel `json:"results"`
}

// Jobs search query
type JobsSearchQueryModel struct {
	ItemsPerPage int `json:"items_per_page"`
	Page         int `json:"page"`
	JobId        int `json:"job_id"`
	JobFilename  string `json:"filename"`
	JobStatus    string `json:"job_status"`
}

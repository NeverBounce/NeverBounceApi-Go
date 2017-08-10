// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// JobsSearchRequestModel is the request model for Jobs.Search()
type JobsSearchRequestModel struct {
	GenericRequestModel
	ItemsPerPage int `json:"items_per_page,omitempty"`
	Page         int `json:"page,omitempty"`
	JobID        int `json:"job_id,omitempty"`
	JobFilename  string `json:"filename,omitempty"`
	JobStatus    string `json:"job_status,omitempty"`
}

// JobsSearchResponseModel is the request model for Jobs.Search()
type JobsSearchResponseModel struct {
	GenericResponseModel
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Query        JobsSearchQueryModel `json:"query"`
	Results      []JobStatusModel `json:"results"`
}

// JobsSearchQueryModel contains the search query from the request
type JobsSearchQueryModel struct {
	ItemsPerPage int `json:"items_per_page"`
	Page         int `json:"page"`
	JobID        int `json:"job_id"`
	JobFilename  string `json:"filename"`
	JobStatus    string `json:"job_status"`
}

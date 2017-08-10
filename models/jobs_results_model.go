// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// JobsResultsRequestModel is the request model for Jobs.Results()
type JobsResultsRequestModel struct {
	GenericRequestModel
	ItemsPerPage int `json:"items_per_page,omitempty"`
	Page         int `json:"page,omitempty"`
	JobID        int `json:"job_id,omitempty"`
}

// JobsResultsResponseModel is the response model for Jobs.Results()
type JobsResultsResponseModel struct {
	GenericResponseModel
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Query        JobsSearchQueryModel `json:"query"`
	Results      []JobsResultsDataModel `json:"results"`
}

// JobsResultsQueryModel contains the query info for the request
type JobsResultsQueryModel struct {
	ItemsPerPage int `json:"items_per_page"`
	Page         int `json:"page"`
	JobID        int `json:"job_id,omitempty"`
}

// JobsResultsDataModel is the model for the each result returned by Jobs.Results()
type JobsResultsDataModel struct {
	Verification VerificationModel `json:"verification"`
	Data         map[string]interface{} `json:"data"`
}

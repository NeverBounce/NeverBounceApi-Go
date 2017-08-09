package nbModels

// Job Search request model
type JobsResultsRequestModel struct {
	GenericRequestModel
	ItemsPerPage int `json:"items_per_page,omitempty"`
	Page         int `json:"page,omitempty"`
	JobId        int `json:"job_id,omitempty"`
}

// Job Model response
type JobsResultsResponseModel struct {
	GenericResponseModel
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Query        JobsSearchQueryModel `json:"query"`
	Results      []JobsResultsDataModel `json:"results"`
}

// Jobs search query
type JobsResultsQueryModel struct {
	ItemsPerPage int `json:"items_per_page"`
	Page         int `json:"page"`
	JobId        int `json:"job_id,omitempty"`
}

type JobsResultsDataModel struct {
	Verification VerificationModel `json:"verification"`
	Data         map[string]interface{} `json:"data"`
}

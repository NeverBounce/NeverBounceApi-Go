// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// JobsStartRequestModel is the request model for Jobs.Start()
type JobsStartRequestModel struct {
	GenericRequestModel
	JobID             int  `json:"job_id"`
	RunSample         bool `json:"run_sample,omitempty"`
	AllowManualReview bool `json:"allow_manual_review,omitempty"`
}

// JobsStartResponseModel is the response model for Jobs.Start()
type JobsStartResponseModel struct {
	GenericResponseModel
	QueueID string `json:"queue_id"`
}

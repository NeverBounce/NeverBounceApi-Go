// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// JobsParseRequestModel is the request model for Jobs.Parse()
type JobsParseRequestModel struct {
	GenericRequestModel
	JobID     int  `json:"job_id"`
	AutoStart bool `json:"auto_start,omitempty"`
}

// JobsParseResponseModel is the response model for Jobs.Parse()
type JobsParseResponseModel struct {
	GenericResponseModel
	QueueID string `json:"queue_id"`
}

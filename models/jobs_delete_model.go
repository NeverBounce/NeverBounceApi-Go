// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// JobsDeleteRequestModel is request model for Jobs.Delete()
type JobsDeleteRequestModel struct {
	GenericRequestModel
	JobID int `json:"job_id"`
}

// JobsDeleteResponseModel is the response model for Jobs.Delete()
type JobsDeleteResponseModel struct {
	GenericResponseModel
}

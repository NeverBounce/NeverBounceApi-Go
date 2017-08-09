package nbModels

type JobsParseRequestModel struct {
	GenericRequestModel
	JobId     int `json:"job_id"`
	AutoStart bool `json:"auto_start,omitempty"`
}

// ParseInfo : Result model of Info API
type JobsParseResponseModel struct {
	GenericResponseModel
	QueueID string `json:"queue_id"`
}

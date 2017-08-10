package nbModels

type JobsStartRequestModel struct {
	GenericRequestModel
	JobId     int `json:"job_id"`
	RunSample bool `json:"run_sample,omitempty"`
}

type JobsStartResponseModel struct {
	GenericResponseModel
	QueueId string `json:"queue_id"`
}

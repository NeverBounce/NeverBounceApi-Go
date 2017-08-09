package nbModels

type JobsDeleteRequestModel struct {
	GenericRequestModel
	JobId     int `json:"job_id"`
}

type JobsDeleteResponseModel struct {
	GenericResponseModel
}

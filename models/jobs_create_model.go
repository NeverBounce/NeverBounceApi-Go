// Package nbDto holds API result models.
package nbModels

// Creation request model
type JobsCreateSuppliedDataRequestModel struct {
	GenericRequestModel
	InputLocation string `json:"input_location"`
	SuppliedData  map[int]interface{} `json:"input"`
	AutoParse     bool `json:"auto_parse"`
	AutoRun       bool `json:"auto_run"`
	RunSample     bool `json:"run_sample"`
	FileName      string `json:"filename,omitempty"`
}

// Creation request model
type JobsCreateRemoteUrlRequestModel struct {
	GenericRequestModel
	InputLocation string `json:"input_location"`
	RemoteUrl     string `json:"input"`
	AutoParse     bool `json:"auto_parse"`
	AutoRun       bool `json:"auto_run"`
	RunSample     bool `json:"run_sample"`
	FileName      string `json:"filename,omitempty"`
}

// Creation response
type JobsCreateResponseModel struct {
	GenericResponseModel
	JobId int `json:"job_id"`
}

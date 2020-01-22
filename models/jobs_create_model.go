// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// JobsCreateSuppliedDataRequestModel is the request model for creating a job with SuppliedData.
// See examples/main.go for an example of it's use
type JobsCreateSuppliedDataRequestModel struct {
	GenericRequestModel
	InputLocation  string              `json:"input_location"`
	SuppliedData   map[int]interface{} `json:"input"`
	AutoParse      bool                `json:"auto_parse"`
	AutoStart      bool                `json:"auto_start"`
	RunSample      bool                `json:"run_sample"`
	FileName       string              `json:"filename,omitempty"`
	HistoricalData HistoricalDataModel `json:"request_meta_data, omitempty"`
}

// JobsCreateRemoteURLRequestModel is the request model for creating a job with a remote URL
type JobsCreateRemoteURLRequestModel struct {
	GenericRequestModel
	InputLocation  string              `json:"input_location"`
	RemoteURL      string              `json:"input"`
	AutoParse      bool                `json:"auto_parse"`
	AutoStart      bool                `json:"auto_start"`
	RunSample      bool                `json:"run_sample"`
	FileName       string              `json:"filename,omitempty"`
	HistoricalData HistoricalDataModel `json:"request_meta_data, omitempty"`
}

// JobsCreateResponseModel is the response model for both creation methods
type JobsCreateResponseModel struct {
	GenericResponseModel
	JobID int `json:"job_id"`
}

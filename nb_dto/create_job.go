// Package nbDto holds API result models.
package nbDto

// CreateJob : Post model of Create job API
type CreateJob struct {
	APIKEY        string `json:"key"`
	InputLocation string `json:"input_location"`
	Input         []string `json:"input"`
	AutoParse     bool `json:"auto_parse"`
	AutoRun       bool `json:"auto_run"`
	RunSample     bool `json:"run_sample"`
	FileName      string `json:"filename"`
}

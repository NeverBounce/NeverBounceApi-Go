// Package nbDto holds API result models.
package nbDto

// CreateSearch : Post model of Create search API
type CreateSearch struct {
	ApiKEY        string `json:"key"`
	InputLocation string `json:"input_location"`
	Input         []string `json:"input"`
	AutoParse     bool `json:"auto_parse"`
	AutoRun       bool `json:"auto_run"`
	RunSample     bool `json:"run_sample"`
	FileName      string `json:"filename"`
}

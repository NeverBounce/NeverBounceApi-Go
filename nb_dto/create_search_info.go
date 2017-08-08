// Package nbDto holds API result models.
package nbModels

// CreateSearchInfo : Result model of Create search API
type CreateSearchInfo struct {
	Status    string `json:"status"`
	JobID     int `json:"job_id"`
	ExecutionTime int `json:"execution_time"`
}

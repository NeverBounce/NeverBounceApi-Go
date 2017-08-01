// Package nbDto holds API result models.
package nbDto

// ParseInfo : Result model of Info API
type ParseInfo struct {
	Status        string `json:"status"`
	QueueID int `json:"queue_id"`
	ExecutionTime int `json:"execution_time"`
}

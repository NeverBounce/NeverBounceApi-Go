// Package nbDto holds API result models.
package nbDto

// StartInfo : Result model of Info API
type StartInfo struct {
	Status        string `json:"status"`
	QueueID int `json:"queue_id"`
	ExecutionTime int `json:"execution_time"`
}

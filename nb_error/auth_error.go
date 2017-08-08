// Package nbError holds API error result models.
package nbErrors

// ApiError : Error Result model of Auth API
type ApiError struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	ExecutionTime int `json:"execution_time"`
}

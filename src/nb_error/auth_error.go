package nbError

type AuthError struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	ExecutionTime int `json:"execution_time"`
}

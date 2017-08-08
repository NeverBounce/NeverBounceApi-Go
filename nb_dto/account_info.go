// Package nbDto holds API result models.
package nbModels

// AccountInfo : Result model of Info API
type AccountInfo struct {
	Status                string `json:"status"`
	BillingType           string `json:"billing_type"`
	Credits               int `json:"credits"`
	FreeAPICredits        int `json:"free_api_credits"`
	MonthlyAPIUsage       int `json:"monthly_api_usage"`
	MonthlyDashboardUsage int `json:"monthly_dashboard_usage"`
	JobsCompleted         int `json:"jobs_completed"`
	JobsUnderReview       int `json:"jobs_under_review"`
	JobsQueued            int `json:"jobs_queued"`
	JobsProcessing        int `json:"jobs_processing"`
	ExecutionTime         int `json:"execution_time"`
}

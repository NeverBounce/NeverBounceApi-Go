// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// AccountInfoResponseModel is used to parse the response from Account.Info()
type AccountInfoResponseModel struct {
	GenericResponseModel
	BillingType string `json:"billing_type"`
	CreditsInfo CreditsInfoModel `json:"credits_info"`
	JobsCounts  JobCountsModel `json:"job_counts"`
}

// JobCountsModel is used to parse the JobsCounts from the AccountInfoResponseModel
type JobCountsModel struct {
	Completed   int `json:"completed"`
	UnderReview int `json:"under_review"`
	Queued      int `json:"queued"`
	Processing  int `json:"processing"`
}

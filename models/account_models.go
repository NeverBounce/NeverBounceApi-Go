// Package nbDto holds API result models.
package nbModels

// AccountInfoModel : Result model of Info API
type AccountInfoResponseModel struct {
	GenericResponseModel
	BillingType string `json:"billing_type"`
	CreditsInfo CreditsInfoModel `json:"credits_info"`
	JobsCounts  JobCountsModel `json:"job_counts"`
}

type JobCountsModel struct {
	Completed   int `json:"completed"`
	UnderReview int `json:"under_review"`
	Queued      int `json:"queued"`
	Processing  int `json:"processing"`
}

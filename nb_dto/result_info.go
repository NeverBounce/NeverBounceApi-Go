// Package nbDto holds API result models.
package nbDto

// ResultQuery : Nested result model of Info API
type ResultQuery struct {
	JobID        int `json:"job_id"`
	Valids       int `json:"valids"`
	InValids     int `json:"invalids"`
	Disposables  int `json:"disposables"`
	Catchalls    int `json:"catchalls"`
	Unknowns     int `json:"unknowns"`
	Page         int `json:"job_id"`
	ItemsPerPage int `json:"items_per_page"`
}

// Data : The data object will contain the original data submitted for this row.
// If the source data was submitted via the API the data will have the same keys as when it was originally submitted.
// If submitted as a CSV in the dashboard the data will use the header row to determine the keys for the data when available.
type Data struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

// Verification : The verification object contains the verification results for the email being verified.
// This object contains the same information that the single/check API returns.
type Verification struct {
	Result              string `json:"result"`
	Flags               []string `json:"flags"`
	SuggestedCorrection string `json:"suggested_correction"`
	AddressInfo         AddressInfo `json:"address_info"`
}

// Result : Result model of Result API
type Result struct {
	Data         Data `json:"data"`
	Verification Verification `json:"verification"`
}

// ResultInfo : Nested result model of Info API
type ResultInfo struct {
	Status        string `json:"status"`
	TotalResults  int `json:"total_results"`
	TotalPages    int `json:"total_pages"`
	Query         Query `json:"query"`
	Results       []Result `json:"results"`
	ExecutionTime int `json:"execution_time"`
}
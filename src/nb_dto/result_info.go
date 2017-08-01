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

// Data : Email data of Result API
type Data struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

// Verification : Email verification data of Result API
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

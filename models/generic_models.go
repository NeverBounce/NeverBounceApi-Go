// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// GenericRequestModel is our most basic request model.
// It accepts the APIKey to make the request from
type GenericRequestModel struct {
	APIKey string `json:"key"`
}

// GenericResponseModel is our most basic API response.
// It contains the responses Status and ExecutionTime
type GenericResponseModel struct {
	Status        string `json:"status"`
	ExecutionTime int    `json:"execution_time"`
}

// CreditsInfoModel is used to parse out the credit info the repsonse
type CreditsInfoModel struct {
	PaidCreditsUsed      int    `json:"paid_credits_used"`
	FreeCreditsUsed      int    `json:"free_credits_used"`
	PaidCreditsRemaining int    `json:"paid_credits_remaining"`
	FreeCreditsRemaining int    `json:"free_credits_remaining"`
	MonthlyAPIUsage      string `json:"monthly_api_usage"`
}

// VerificationModel is used to parse out the verification responses from JobsStatus and SingleCheck
type VerificationModel struct {
	Result              string           `json:"result"`
	Flags               []string         `json:"flags"`
	SuggestedCorrection string           `json:"suggested_correction"`
	RetryToken          string           `json:"retry_token"`
	AddressInfo         AddressInfoModel `json:"address_info"`
}

// AddressInfoModel is used to parse out address info from the VerificationModel
type AddressInfoModel struct {
	OriginalEmail   string `json:"original_email"`
	NormalizedEmail string `json:"normalized_email"`
	Addr            string `json:"addr"`
	Alias           string `json:"alias"`
	Host            string `json:"host"`
	Fqdn            string `json:"fqdn"`
	Domain          string `json:"domain"`
	Subdomain       string `json:"subdomain"`
	Tld             string `json:"tld"`
}

// APIErrorModel is our generic error message response
type APIErrorModel struct {
	GenericResponseModel
	Message string `json:"message"`
}

// HistoricalDataModel is used in some requests to enable or disable Historical Data
type HistoricalDataModel struct {
	RequestMetaData int `json:"leverage_historical_data,omitempty"`
}

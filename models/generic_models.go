/*
Package nbModels provides the data mappings for API requests and API responses
*/
package nbModels

// GenericRequestModel: Our generic request struct, expects the apikey
type GenericRequestModel struct {
	ApiKey string `json:"key"`
}

// GenericResponseModel: Our generic response struct, expects the status and execution_time
type GenericResponseModel struct {
	Status        string `json:"status"`
	ExecutionTime int `json:"execution_time"`
}

// Generic credit info struct
type CreditsInfoModel struct {
	PaidCreditsUsed      int `json:"paid_credits_used"`
	FreeCreditsUsed      int `json:"free_credits_used"`
	PaidCreditsRemaining int `json:"paid_credits_remaining"`
	FreeCreditsRemaining int `json:"free_credits_remaining"`
	MonthlyAPIUsage      string `json:"monthly_api_usage"`
}

// Generic Verification struct
type VerificationModel struct {
	Result              string `json:"result"`
	Flags               []string `json:"flags"`
	SuggestedCorrection string `json:"suggested_correction"`
	RetryToken          string `json:"retry_token"`
	AddressInfo         AddressInfoModel `json:"address_info"`
}

// Generic Address info struct
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

type ApiErrorModel struct {
	GenericResponseModel
	Message string `json:"message"`
}
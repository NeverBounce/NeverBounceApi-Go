package nbDto

type AddressInfo struct {
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

type CreditsInfo struct {
	PaidCreditsUsed      int `json:"paid_credits_used"`
	FreeCreditsUsed      int `json:"free_credits_used"`
	PaidCreditsRemaining int `json:"free_credits_remaining"`
	FreeCreditsRemaining int `json:"free_credits_remaining"`
	MonthlyApiUsage      string `json:"monthly_api_usage"`
}

type SingleCheckInfo struct {
	Status                 string `json:"status"`
	Result                 string `json:"result"`
	Flags                  []string `json:"flags"`
	SuggestedCorrection    string `json:"suggested_correction"`
	RetryToken             string `json:"retry_token"`
	ExecutionTime          int `json:"execution_time"`
	AddressInfo            AddressInfo `json:"address_info"`
	CreditsInfo CreditsInfo `json:"credits_info"`
}

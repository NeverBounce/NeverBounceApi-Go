// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// SingleCheckRequestModel is the request model for Single.Check()
type SingleCheckRequestModel struct {
	GenericRequestModel
	Email       string `json:"email"`
	AddressInfo bool   `json:"address_info,omitempty"`
	CreditInfo  bool   `json:"credits_info,omitempty"`
	Timeout     int    `json:"timeout,omitempty"`
}

// SingleCheckResponseModel is the response model for Single.Check()
type SingleCheckResponseModel struct {
	GenericResponseModel
	VerificationModel
	CreditsInfo CreditsInfoModel `json:"credits_info"`
}

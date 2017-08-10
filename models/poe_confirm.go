// Package nbModels provides the data mappings for API requests and API responses
package nbModels

// POEConfirmRequestModel is the request model for POE.Confirm()
type POEConfirmRequestModel struct {
	GenericRequestModel
	Email             string `json:"email"`
	ConfirmationToken string `json:"confirmation_token"`
	TransactionID     string `json:"transaction_id"`
	Result            string `json:"result"`
}

// POEConfirmResponseModel is the response model for POE.Confirm()
type POEConfirmResponseModel struct {
	GenericResponseModel
	Confirmed bool `json:"token_confirmed"`
}

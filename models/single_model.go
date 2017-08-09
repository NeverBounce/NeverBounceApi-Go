package nbModels

// Result model of Single check API
type SingleCheckModel struct {
	GenericResponseModel
	VerificationModel
	CreditsInfo         CreditsInfoModel `json:"credits_info"`
}
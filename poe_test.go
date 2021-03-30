package neverbounce_test

import (
	neverbounce "github.com/NeverBounce/NeverBounceApi-Go"
	nbModels "github.com/NeverBounce/NeverBounceApi-Go/models"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NeverBounce", func() {
	Describe("POE", func() {
		It("should return an instance of POEConfirmResponseModel", func() {
			response := httpmock.NewStringResponse(200, `{
				"status": "success",
				"token_confirmed": true,
				"execution_time": 300
			}`)
			response.Header.Set("content-type", "application/json")
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/poe/confirm",
				httpmock.ResponderFromResponse(response))

			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.POE.Confirm(&nbModels.POEConfirmRequestModel{
				Email:             "support@neverbounce.com",
				ConfirmationToken: "1341234jkh12h34lb2134b143",
				TransactionID:     "1340813265013984123",
				Result:            "valid",
			})
			Expect(resp).NotTo(BeNil())
			Expect(resp.Confirmed).To(Equal(true))
			Expect(err).To(BeNil())
		})
	})
})

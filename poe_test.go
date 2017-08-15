package neverbounce_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/NeverBounce/NeverBounceApi-Go"
	"gopkg.in/jarcoal/httpmock.v1"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

var _ = Describe("NeverBounce", func() {
	Describe("POE", func() {
		It("should return an instance of POEConfirmResponseModel", func() {
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/poe/confirm",
				httpmock.NewStringResponder(200, `{
					"status": "success",
					"token_confirmed": true,
					"execution_time": 300
				}`))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.POE.Confirm(&nbModels.POEConfirmRequestModel{
				Email: "support@neverbounce.com",
				ConfirmationToken: "1341234jkh12h34lb2134b143",
				TransactionID: "1340813265013984123",
				Result: "valid",
			})
			Expect(resp).NotTo(BeNil())
			Expect(resp.Confirmed).To(Equal(true))
			Expect(err).To(BeNil())
		})
	})
})

package neverbounce_test

import (
	"github.com/NeverBounce/NeverBounceApi-Go"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("NeverBounce", func() {
	Describe("Check", func() {
		It("should return an instance of VerificationObject with a good response and error should be nil", func() {
			response := httpmock.NewStringResponse(200, `{
				"status": "success",
				"result": "valid",
				"flags": [
				"has_dns",
				"has_dns_mx",
				"role_account"
				],
				"suggested_correction": "",
				"retry_token": "",
				"execution_time": 399
			}`)
			response.Header.Set("content-type", "application/json")
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/single/check",
				httpmock.ResponderFromResponse(response))

			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Single.Check(&nbModels.SingleCheckRequestModel{
				Email: "support@neverbounce.com",
			})
			Expect(resp).NotTo(BeNil())
			Expect(resp.Result).To(Equal("valid"))
			Expect(err).To(BeNil())
		})
	})
})

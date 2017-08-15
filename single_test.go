package neverbounce_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/NeverBounce/NeverBounceApi-Go"
	"gopkg.in/jarcoal/httpmock.v1"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

var _ = Describe("NeverBounce", func() {
	Describe("Check", func() {
		It("should return an instance of VerificationObject with a good response and error should be nil", func() {
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/single/check",
				httpmock.NewStringResponder(200, `{
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
				}`))
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

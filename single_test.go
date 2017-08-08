package neverbounce_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/NeverBounce/NeverBounceApi-Go"
	"gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("NeverBounce", func() {
	Describe("Check", func() {
		It("should return an instance of VerificationObject with a good response and error should be nil", func() {
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/single/check?key=apiKey&email=support@neverbounce.com",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "result": "valid",
                "flags": [
                    "has_dns",
                    "has_dns_mx"
                ],
                "suggested_correction": "",
                "retry_token": "",
                "execution_time": 499
            }`))
			neverBounce, _ := neverbounce.New("apiKey")
			resp, err := neverBounce.Single.Check("support@neverbounce.com", false, false, "")
			Expect(resp).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})

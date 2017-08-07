package neverBounce_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/jarcoal/httpmock.v1"
	. "github.com/onsi/gomega"
	"github.com/NeverBounce/NeverBounceApi-Go"
)

var _ = Describe("Account", func() {
	Describe("Info", func() {
		It("should return a object during a good response", func() {
			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/account/info?key=apiKey",
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
			neverBounce, err := neverBounce.New("apiKey")
			Expect(neverBounce).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})

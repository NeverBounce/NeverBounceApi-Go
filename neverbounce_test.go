package neverbounce_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/NeverBounce/NeverBounceApi-Go"
	"testing"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestNeverBounceApiGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NeverBounceApiGo Suite")
}

var _ = Describe("NeverBounce", func() {
	Describe("MakeRequest", func() {
		It("should return an error during a 404", func() {
			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/account/info",
				httpmock.NewStringResponder(404, `404 Not Found`))
			neverBounce, _ := neverbounce.New("apiKey")
			resp, err := neverBounce.Account.Info()
			Expect(resp).To(BeNil())
			Expect(err).NotTo(BeNil())
		})

		It("should return an error during a 503", func() {
			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/account/info",
				httpmock.NewStringResponder(503, `503 `))
			neverBounce, _ := neverbounce.New("apiKey")
			resp, err := neverBounce.Account.Info()
			Expect(resp).To(BeNil())
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("We were unable to complete your request. " +
				"The following information was supplied: 503" +
				"\n\n(Internal error [status  503])"))
		})

		It("should return an error when status isn't 'success'", func() {
			response := httpmock.NewStringResponse(200, `{
					"status": "general_failure",
					"message": "Something went wrong",
					"execution_time": 300
				}`)
			response.Header.Set("Content-Type", "application/json")

			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/account/info",
				httpmock.ResponderFromResponse(response))
			neverBounce, _ := neverbounce.New("apiKey")
			body, err := neverBounce.Account.Info()
			Expect(body).To(BeNil())
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("We were unable to complete your request. " +
				"The following information was supplied: Something went wrong" +
				"\n\n(general_failure)"))
		})
	})
})

var _ = BeforeSuite(func() {
	// block all HTTP requests
	httpmock.Activate()
	// mock the root info API
	httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/account/info",
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
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})

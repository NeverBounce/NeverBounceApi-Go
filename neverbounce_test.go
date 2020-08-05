package neverbounce_test

import (
	"github.com/NeverBounce/NeverBounceApi-Go"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/jarcoal/httpmock.v1"
	"os"
	"testing"
)

func TestNeverBounceApiGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NeverBounceApiGo Suite")
}

var _ = Describe("NeverBounce", func() {
	Describe("MakeRequest", func() {

		It("should return an error when the content-type mismatches", func() {
			response := httpmock.NewStringResponse(200, ``)
			response.Header.Set("content-type", "text/html")

			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/account/info",
				httpmock.ResponderFromResponse(response))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Account.Info()
			Expect(resp).To(BeNil())
			Expect(err).NotTo(BeNil())
			if nbError, ok := err.(*neverbounce.Error); ok {
				Expect(nbError.Message).To(Equal("The API responded with a datatype of \"text/html" +
					"\", but \"application/json\" was expected." +
					"\n\n(Internal error [status 200])"))
			}
		})

		It("should not return an error when requesting job/download with non json response", func() {
			response := httpmock.NewStringResponse(200, ``)
			response.Header.Set("content-type", "text/html")

			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/jobs/download",
				httpmock.ResponderFromResponse(response))
			neverBounce := neverbounce.New("apiKey")
			err := neverBounce.Jobs.Download(&nbModels.JobsDownloadRequestModel{
				JobID:            296050,
				EmailStatusAsInt: true,
			}, "test.csv")
			Expect(err).To(BeNil())

			err2 := os.Remove("test.csv")
			if err2 != nil {
				panic(err2)
			}
		})

		It("should return an error when the response is empty and content-type matches", func() {
			response := httpmock.NewStringResponse(200, ``)
			response.Header.Set("content-type", "application/json")

			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/account/info",
				httpmock.ResponderFromResponse(response))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Account.Info()
			Expect(resp).To(BeNil())
			Expect(err).NotTo(BeNil())
			if nbError, ok := err.(*neverbounce.Error); ok {
				Expect(nbError.Message).To(Equal("We were unable to parse the API response. " +
					"The following information was supplied: " +
					"\n\n(Internal error [status 200])"))
			}
		})

		It("should return an error during a 404", func() {
			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/account/info",
				httpmock.NewStringResponder(404, `404 Not Found`))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Account.Info()
			Expect(resp).To(BeNil())
			Expect(err).NotTo(BeNil())
			if nbError, ok := err.(*neverbounce.Error); ok {
				Expect(nbError.Message).To(Equal("We were unable to complete your request. " +
					"The following information was supplied: 404 Not Found" +
					"\n\n(Request error [status 404])"))
			}
		})

		It("should return an error during a 503", func() {
			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/account/info",
				httpmock.NewStringResponder(503, `503 Temporarily Unavailable`))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Account.Info()
			Expect(resp).To(BeNil())
			Expect(err).NotTo(BeNil())
			if nbError, ok := err.(*neverbounce.Error); ok {
				Expect(nbError.Message).To(Equal("We were unable to complete your request. " +
					"The following information was supplied: 503 Temporarily Unavailable" +
					"\n\n(Internal error [status 503])"))
			}
		})

		It("should return an error when status isn't 'success'", func() {
			response := httpmock.NewStringResponse(200, `{
					"status": "general_failure",
					"message": "Something went wrong",
					"execution_time": 300
				}`)
			response.Header.Set("content-Type", "application/json")

			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/account/info",
				httpmock.ResponderFromResponse(response))
			neverBounce := neverbounce.New("apiKey")
			body, err := neverBounce.Account.Info()
			Expect(body).To(BeNil())
			Expect(err).NotTo(BeNil())
			if nbError, ok := err.(*neverbounce.Error); ok {
				Expect(nbError.Message).To(Equal("We were unable to complete your request. " +
					"The following information was supplied: Something went wrong" +
					"\n\n(general_failure)"))
			}
		})
	})
})

var _ = BeforeSuite(func() {
	// block all HTTP requests
	httpmock.Activate()
	// mock the root info API
	httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4.2/account/info",
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

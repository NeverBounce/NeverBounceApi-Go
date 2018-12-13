package neverbounce_test

import (
	"github.com/NeverBounce/NeverBounceApi-Go"
	"gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("Account", func() {
	Describe("Info", func() {
		It("should return a object during a good response and error should be nil", func() {
			// mock the root info API
			response := httpmock.NewStringResponse(200, `{
				"status": "success",
				"credits_info": {
					"paid_credits_used": 0,
					"free_credits_used": 0,
					"paid_credits_remaining": 9950791,
					"free_credits_remaining": 0
				},
				"job_counts": {
					"completed": 409,
					"under_review": 0,
					"queued": 0,
					"processing": 0
				},
				"execution_time": 896
			}`)
			response.Header.Set("content-type", "application/json")
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/account/info",
				httpmock.ResponderFromResponse(response))

			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Account.Info()
			Expect(resp).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})

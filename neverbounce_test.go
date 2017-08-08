package neverbounce_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestNeverBounceApiGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NeverBounceApiGo Suite")
}
var _ = BeforeSuite(func() {
	// block all HTTP requests
	httpmock.Activate()
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
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})
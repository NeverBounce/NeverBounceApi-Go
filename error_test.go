package neverbounce_test

import (
	"github.com/NeverBounce/NeverBounceApi-Go"
)

var _ = Describe("Jobs", func() {

	It("should expand error strings", func() {
		var err = neverbounce.Error{
			Type:    "auth_failure",
			Message: "We were unable to authenticate your request. The following information was supplied: Invalid API key 'api_key'\n\n(auth_failure)",
		}
		Expect(err.Error()).To(ContainSubstring("{\"status\":\"auth_failure\",\"message\":\"We were unable to authenticate your request. The following information was supplied: Invalid API key 'api_key'"))
	})
})

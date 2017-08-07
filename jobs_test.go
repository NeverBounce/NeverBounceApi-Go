package neverBounce_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/jarcoal/httpmock.v1"
	. "github.com/onsi/gomega"
	"github.com/NeverBounce/NeverBounceApi-Go"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_dto"
	"fmt"
)

var _ = Describe("Jobs", func() {
	Describe("Create", func() {
		It("should return JobID not equal to zero and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("POST", "https://api.neverbounce.com/v4/jobs/create",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "job_id": 150970,
                "execution_time": 388
            }`))
			neverBounce, _ := neverBounce.New("apiKey")
			resp, err := neverBounce.Jobs.Create(&nbDto.CreateSearch{
				InputLocation: "supplied",
				Input:         []string{"enkhalifapro@gmail.com"},
				AutoParse:     true,
				AutoRun:       true,
				RunSample:     false,
				FileName:      "example.csv"})
			fmt.Println("xxxxxx")
			fmt.Println(resp.JobID)
			Expect(resp.JobID).NotTo(BeZero())
			Expect(err).To(BeNil())
		})
	})
})

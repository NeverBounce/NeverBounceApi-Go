package neverBounce_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/jarcoal/httpmock.v1"
	. "github.com/onsi/gomega"
	"github.com/NeverBounce/NeverBounceApi-Go"
	"github.com/NeverBounce/NeverBounceApi-Go/nb_dto"
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
			Expect(resp.JobID).NotTo(BeZero())
			Expect(err).To(BeNil())
		})
	})
	Describe("Parse", func() {
		It("should return a valid queueID and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("POST", "https://api.neverbounce.com/v4/jobs/parse",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "queue_id": 55,
                "execution_time": 388
            }`))
			neverBounce, _ := neverBounce.New("apiKey")
			resp, err := neverBounce.Jobs.Parse(150970, false)
			Expect(resp.QueueID).To(Equal(55))
			Expect(err).To(BeNil())
		})
	})
	Describe("Start", func() {
		It("should return a valid queueID and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("POST", "https://api.neverbounce.com/v4/jobs/start",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "queue_id": 55,
                "execution_time": 388
            }`))
			neverBounce, _ := neverBounce.New("apiKey")
			resp, err := neverBounce.Jobs.Start(150970, false)
			Expect(resp.QueueID).To(Equal(55))
			Expect(err).To(BeNil())
		})
	})
})

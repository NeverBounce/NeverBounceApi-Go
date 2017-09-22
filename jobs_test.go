package neverbounce_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/jarcoal/httpmock.v1"
	. "github.com/onsi/gomega"
	"github.com/NeverBounce/NeverBounceApi-Go"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
	"os"
)

var _ = Describe("Jobs", func() {
	Describe("CreateFromSuppliedData", func() {
		It("should return a JobID not equal to zero and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("POST", "https://api.neverbounce.com/v4/jobs/create",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "job_id": 150970,
                "execution_time": 388
            }`))

			// Build data map
			createData := map[int]interface{}{}
			createData[0] = map[string]interface{}{
				"id":    12345,
				"email": "support@neverbounce.com",
				"name":  "Bob McValid",
			}
			createData[1] = map[string]interface{}{
				"id":    12346,
				"email": "invalid@neverbounce.com",
				"name":  "Fred McInvalid",
			}

			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Jobs.CreateFromSuppliedData(&nbModels.JobsCreateSuppliedDataRequestModel{
				InputLocation: "supplied",
				SuppliedData:  createData,
				AutoParse:     true,
				AutoStart:     true,
				RunSample:     false,
				FileName:      "example.csv"})
			Expect(resp.JobID).To(Equal(150970))
			Expect(err).To(BeNil())
		})
	})

	Describe("CreateFromRemoteURL", func() {
		It("should return a JobID not equal to zero and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("POST", "https://api.neverbounce.com/v4/jobs/create",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "job_id": 150970,
                "execution_time": 388
            }`))

			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Jobs.CreateFromRemoteURL(&nbModels.JobsCreateRemoteURLRequestModel{
				InputLocation: "supplied",
				RemoteURL:     "https://example.com/file.csv",
				AutoParse:     true,
				AutoStart:     true,
				RunSample:     false,
				FileName:      "example.csv"})
			Expect(resp.JobID).To(Equal(150970))
			Expect(err).To(BeNil())
		})
	})

	Describe("Parse", func() {
		It("should return a valid QueueID and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("POST", "https://api.neverbounce.com/v4/jobs/parse",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "queue_id": "NB-PQ-59246392E9E5D",
                "execution_time": 388
            }`))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Jobs.Parse(&nbModels.JobsParseRequestModel{
				JobID: 150970,
			})
			Expect(resp.QueueID).To(Equal("NB-PQ-59246392E9E5D"))
			Expect(err).To(BeNil())
		})
	})

	Describe("Start", func() {
		It("should return a valid QueueID and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("POST", "https://api.neverbounce.com/v4/jobs/start",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "queue_id": "NB-PQ-59246392E9E5D",
                "execution_time": 388
            }`))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Jobs.Start(&nbModels.JobsStartRequestModel{
				JobID: 150970,
			})
			Expect(resp.QueueID).To(Equal("NB-PQ-59246392E9E5D"))
			Expect(err).To(BeNil())
		})
	})

	Describe("Status", func() {
		It("should return a valid TotalRecords and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/jobs/status",
				httpmock.NewStringResponder(200, `{
					"status": "success",
					"id": 277461,
					"filename": "Created from Array.csv",
					"created_at": "2017-07-25 14:52:27",
					"started_at": "2017-07-25 14:52:40",
					"finished_at": "2017-07-25 14:53:06",
					"total": {
						"records": 2,
						"billable": 2,
						"processed": 2,
						"valid": 0,
						"invalid": 2,
						"catchall": 0,
						"disposable": 0,
						"unknown": 0,
						"duplicates": 0,
						"bad_syntax": 0
					},
					"bounce_estimate": 0,
					"percent_complete": 100,
					"job_status": "complete",
					"execution_time": 322
					}`))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Jobs.Status(&nbModels.JobsStatusRequestModel{
				JobID: 277461,
			})
			Expect(resp.JobID).To(Equal(277461))
			Expect(resp.Totals.Records).To(Equal(2))
			Expect(resp.JobStatus).To(Equal("complete"))
			Expect(err).To(BeNil())
		})
	})

	Describe("Result", func() {
		It("should return a valid TotalResults and error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/jobs/results",
				httpmock.NewStringResponder(200, `{
					"status": "success",
					"total_results": 2,
					"total_pages": 1,
					"query": {
						"job_id": 251319,
						"valids": 1,
						"invalids": 1,
						"disposables": 1,
						"catchalls": 1,
						"unknowns": 1,
						"page": 0,
						"items_per_page": 10
					},
					"results": [
						{
						  "data": {
							"email": "support@neverbounce.com",
							"id": "12345",
							"name": "Fred McValid"
						  },
						  "verification": {
							"result": "valid",
							"flags": [
							  "smtp_connectable",
							  "has_dns",
							  "has_dns_mx",
							  "role_account"
							],
							"suggested_correction": "",
							"address_info": {
							  "original_email": "support@neverbounce.com",
							  "normalized_email": "support@neverbounce.com",
							  "addr": "support",
							  "alias": "",
							  "host": "neverbounce.com",
							  "fqdn": "neverbounce.com",
							  "domain": "neverbounce",
							  "subdomain": "",
							  "tld": "com"
							}
						  }
						},
						{
						  "data": {
							"email": "invalid@neverbounce.com",
							"id": "12346",
							"name": "Bob McInvalid"
						  },
						  "verification": {
							"result": "invalid",
							"flags": [
							  "smtp_connectable",
							  "has_dns",
							  "has_dns_mx"
							],
							"suggested_correction": "",
							"address_info": {
							  "original_email": "invalid@neverbounce.com",
							  "normalized_email": "invalid@neverbounce.com",
							  "addr": "invalid",
							  "alias": "",
							  "host": "neverbounce.com",
							  "fqdn": "neverbounce.com",
							  "domain": "neverbounce",
							  "subdomain": "",
							  "tld": "com"
							}
						  }
						}
					],
					"execution_time": 55
					}`))
			neverBounce := neverbounce.New("apiKey")
			resp, err := neverBounce.Jobs.Results(&nbModels.JobsResultsRequestModel{
				JobID: 251319,
			})
			Expect(resp.TotalResults).To(Equal(2))
			Expect(resp.Results[0].Verification.Result).To(Equal("valid"))
			Expect(err).To(BeNil())
		})
	})

	Describe("Download", func() {
		It("error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("GET", "https://api.neverbounce.com/v4/jobs/download",
				httpmock.NewStringResponder(200, `{
                "status": "success"}`))
			neverBounce := neverbounce.New("apiKey")
			err := neverBounce.Jobs.Download(&nbModels.JobsDownloadRequestModel{
				JobID: 150970,
			}, "./example.csv")

			Expect(err).To(BeNil())
			// Cleanup
			os.Remove("./example.csv")
		})
	})

	Describe("Delete", func() {
		It("error should be nil", func() {
			// mock the root info API
			httpmock.RegisterResponder("POST", "https://api.neverbounce.com/v4/jobs/delete",
				httpmock.NewStringResponder(200, `{
                "status": "success",
                "execution_time": 388
            }`))
			neverBounce := neverbounce.New("apiKey")
			_, err := neverBounce.Jobs.Delete(&nbModels.JobsDeleteRequestModel{
				JobID: 150970,
			})
			Expect(err).To(BeNil())
		})
	})

})

/*
Package main offers several examples for working with the neverbounce package.

To use:
Add your api key and uncomment specific methods to test them out.
*/
package main

import (
	"fmt"
	"github.com/NeverBounce/NeverBounceApi-Go"
	"github.com/NeverBounce/NeverBounceApi-Go/models"
)

func main() {
	// instantiate neverBounce
	client, err := neverbounce.New("secret_nvrbnc_golang_")
	if err != nil {
		panic(err)
	}

	//AccountInfo(client)
	//SingleCheck(client)
	//JobsSearch(client)
	//JobsCreateFromSuppliedData(client)
	//JobsCreateFromRemoteURL(client)
	//JobsParse(client)
	//JobsStart(client)
	//JobsStatus(client)
	//JobsResults(client)
	//JobsDownload(client)
	//JobsDelete(client)
	POEConfirm(client)
}

// AccountInfo demonstrates how to retrieve the account info
func AccountInfo(client *neverbounce.NeverBounce) {
	accountInfo, err := client.Account.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(accountInfo)
}

// SingleCheck demonstrates how to make to verify a single email
func SingleCheck(client *neverbounce.NeverBounce) {
	singleResults, err := client.Single.Check(&nbModels.SingleCheckRequestModel{
		Email:       "support@neverbounce.com",
		AddressInfo: true,
		CreditInfo:  true,
		Timeout:     10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(singleResults)
}

// JobsSearch demonstrates how to search for existing jobs on your account
func JobsSearch(client *neverbounce.NeverBounce) {
	searchResults, err := client.Jobs.Search(&nbModels.JobsSearchRequestModel{
		JobStatus: "complete",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(searchResults)
}

// JobsCreateFromSuppliedData demonstrates how to create a job using data you have available
func JobsCreateFromSuppliedData(client *neverbounce.NeverBounce) {
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

	// Create a job from supplied data
	jobInfo, err := client.Jobs.CreateFromSuppliedData(&nbModels.JobsCreateSuppliedDataRequestModel{
		SuppliedData: createData,
		AutoParse:    false,
		AutoRun:      false,
		RunSample:    false,
		FileName:     "Created from Golang.csv"})
	if err != nil {
		panic(err)
	}
	fmt.Println(jobInfo)
}

// JobsCreateFromRemoteURL demonstrates how to create a job using data hosted on a remote url
func JobsCreateFromRemoteURL(client *neverbounce.NeverBounce) {
	jobInfo, err := client.Jobs.CreateFromRemoteURL(&nbModels.JobsCreateRemoteURLRequestModel{
		RemoteURL: "https://example.com/file.csv",
		AutoParse: true,
		AutoRun:   false,
		RunSample: false,
		FileName:  "Created from Golang.csv"})
	if err != nil {
		panic(err)
	}
	fmt.Println(jobInfo)
}

// JobsParse demonstrates how to parse a job after it's been created with AutoParse set to false
func JobsParse(client *neverbounce.NeverBounce) {
	parseInfo, err := client.Jobs.Parse(&nbModels.JobsParseRequestModel{
		JobID: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(parseInfo)
}

// JobsStart demonstrates how to start a job after it's been parsed and AutoStart was set to false
func JobsStart(client *neverbounce.NeverBounce) {
	startInfo, err := client.Jobs.Start(&nbModels.JobsStartRequestModel{
		JobID: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(startInfo)
}

// JobsStatus demonstrates how to get the jobs status and stats
func JobsStatus(client *neverbounce.NeverBounce) {
	statusInfo, err := client.Jobs.Status(&nbModels.JobsStatusRequestModel{
		JobID: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(statusInfo)
}

// JobsResults demonstrates how to get the jobs results once it has completed verification
func JobsResults(client *neverbounce.NeverBounce) {
	resultsInfo, err := client.Jobs.Results(&nbModels.JobsResultsRequestModel{
		JobID: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resultsInfo)
}

// JobsDownload demonstrates how to download the job results as a CSV file once it has completed verification
func JobsDownload(client *neverbounce.NeverBounce) {
	err := client.Jobs.Download(&nbModels.JobsDownloadRequestModel{
		JobID:            296050,
		EmailStatusAsInt: true,
	}, "test.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data saved to test.csv")
}

// JobsDelete demonstrates how to delete a job
func JobsDelete(client *neverbounce.NeverBounce) {
	deleteInfo, err := client.Jobs.Delete(&nbModels.JobsDeleteRequestModel{
		JobID: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteInfo)
}

// POEConfirm demonstrates how to delete a job
func POEConfirm(client *neverbounce.NeverBounce) {
	confirmInfo, err := client.POE.Confirm(&nbModels.POEConfirmRequestModel{
		Email: "support@neverbounce.com",
		ConfirmationToken: "1341234jkh12h34lb2134b143",
		TransactionID: "1340813265013984123",
		Result: "valid",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(confirmInfo)
}

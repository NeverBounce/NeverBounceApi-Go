/*
Package main offers several examples for working with the neverbounce package.
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
	client, err := neverbounce.New("api_key")
	if err != nil {
		panic(err)
	}

	account_info(client)
	//single_check(client)
	//jobs_search(client)
	//jobs_create_supplied_data(client)
	//jobs_create_remote_url(client)
	//jobs_parse(client)
	//jobs_start(client)
	//jobs_status(client)
	//jobs_results(client)
	//jobs_download(client)
	//jobs_delete(client)
}

func account_info(client *neverbounce.NeverBounce) {
	accountInfo, err := client.Account.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(accountInfo)
}

func single_check(client *neverbounce.NeverBounce) {
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

func jobs_search(client *neverbounce.NeverBounce) {
	searchResults, err := client.Jobs.Search(&nbModels.JobsSearchRequestModel{
		JobStatus: "complete",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(searchResults)
}

func jobs_create_supplied_data(client *neverbounce.NeverBounce) {
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

func jobs_create_remote_url(client *neverbounce.NeverBounce) {
	jobInfo, err := client.Jobs.CreateFromRemoteUrl(&nbModels.JobsCreateRemoteUrlRequestModel{
		RemoteUrl: "https://example.com/file.csv",
		AutoParse: true,
		AutoRun:   false,
		RunSample: false,
		FileName:  "Created from Golang.csv"})
	if err != nil {
		panic(err)
	}
	fmt.Println(jobInfo)
}

func jobs_parse(client *neverbounce.NeverBounce) {
	parseInfo, err := client.Jobs.Parse(&nbModels.JobsParseRequestModel{
		JobId: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(parseInfo)
}

func jobs_start(client *neverbounce.NeverBounce) {
	startInfo, err := client.Jobs.Start(&nbModels.JobsStartRequestModel{
		JobId: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(startInfo)
}

func jobs_status(client *neverbounce.NeverBounce) {
	statusInfo, err := client.Jobs.Status(&nbModels.JobsStatusRequestModel{
		JobId: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(statusInfo)
}

func jobs_results(client *neverbounce.NeverBounce) {
	resultsInfo, err := client.Jobs.Results(&nbModels.JobsResultsRequestModel{
		JobId: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resultsInfo)
}

func jobs_download(client *neverbounce.NeverBounce) {
	err := client.Jobs.Download(&nbModels.JobsDownloadRequestModel{
		JobId:            296050,
		EmailStatusAsInt: true,
	}, "test.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data saved to test.csv")
}

func jobs_delete(client *neverbounce.NeverBounce) {
	deleteInfo, err := client.Jobs.Delete(&nbModels.JobsDeleteRequestModel{
		JobId: 296050,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteInfo)
}

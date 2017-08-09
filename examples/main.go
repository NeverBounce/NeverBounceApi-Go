// Package main show examples of using NeverBounce package
package main

import (
	"fmt"
	//"github.com/NeverBounce/NeverBounceApi-Go/models"
	"github.com/NeverBounce/NeverBounceApi-Go"
)

func main() {
	// instantiate neverBounce
	neverBounce, err := neverbounce.New("secret_nvrbnc_golang_")
	if err != nil {
		panic(err)
	}

	// Info API
	//info, err := neverBounce.Info()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(info)

	//// Single check API
	singleCheckInfo, err := neverBounce.Single.Check("support@neverbounce.com", true, true, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(singleCheckInfo)

	//// Search API
	//
	//searchInfo, err := neverBounce.Jobs.Search(277184, "example.csv", false, false, false, false, false, false, 1, 10)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(searchInfo)
	//
	//// Create search API
	//
	//createSearchInfo, err := neverBounce.Jobs.Create(&nbModels.CreateJob{
	//	InputLocation: "supplied",
	//	Input:         []string{"enkhalifapro@gmail.com"},
	//	AutoParse:     true,
	//	AutoRun:       true,
	//	RunSample:     false,
	//	FileName:      "ayman.csv"})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(createSearchInfo)
	//
	//// Parse job API
	//
	//parseInfo, err := neverBounce.Jobs.Parse(277184, true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(parseInfo)
	//
	//// Start job API
	//
	//startInfo, err := neverBounce.Jobs.Start(277184, true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(startInfo)
	//
	//// job status API
	//
	//statusInfo, err := neverBounce.Jobs.Status(277184)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(statusInfo)
	//
	//// job results API
	//
	//resultsInfo, err := neverBounce.Jobs.Results(277184, 1, 10)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resultsInfo)
	//
	//// job Download API
	//
	//err = neverBounce.Jobs.Download(277184, "./job.csv")
	//if err != nil {
	//	panic(err)
	//}
	//
	//// job Delete API
	//
	//err = neverBounce.Jobs.Delete(277184)
	//if err != nil {
	//	panic(err)
	//}
}

// Package main show examples of using NeverBounce package
package main

import (
	"fmt"
	"github.com/NeverBounce/NeverBounceApi-Go/src"
	"github.com/NeverBounce/NeverBounceApi-Go/src/nb_dto"
)

func main() {
	// instantiate neverBounce
	neverBounce, err := neverBounce.New("secret_nvrbnc_golang")

	if err != nil {
		panic(err)
	}

	// Info API
	info, err := neverBounce.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	// Single check API
	singleCheckInfo, err := neverBounce.Single.Check("enkhalifapro@gmail.com", true, true, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(singleCheckInfo)

	// Create search API

	createSearchInfo, err := neverBounce.Jobs.Create(&nbDto.CreateSearch{
		InputLocation: "supplied",
		Input:         []string{"enkhalifapro@gmail.com"},
		AutoParse:     true,
		AutoRun:       true,
		RunSample:     false,
		FileName:      "ayman.csv"})
	if err != nil {
		panic(err)
	}
	fmt.Println(createSearchInfo)

	// Parse job API

	parseInfo, err := neverBounce.Jobs.Parse(277184, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(parseInfo)

	// Start job API

	startInfo, err := neverBounce.Jobs.Start(277184, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(startInfo)

}

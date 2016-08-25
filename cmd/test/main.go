package main

import (
	"fmt"

	"github.com/mhenderson-so/azure-ea-billing"
)

var (
	//BaseURI is the base of the URI that will be used for accessing the Azure EA Billing portal
	BaseURI = ""
)

func main() {
	eab := azureeabilling.Config{
		EA:     12345678,
		APIKey: "abcdefg",
	}

	resp, err := eab.GetUsageReports()
	if err != nil {
		fmt.Println(err)
		return
	}
	reports := eab.GetMonthReportsCSV(resp.AvailableMonths[20], azureeabilling.DownloadForStructs)
	fmt.Println(reports.SummaryReport)

	structs, err := reports.ConvertToStructs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, l := range structs.DetailReport {
		fmt.Println(*l)
	}
}

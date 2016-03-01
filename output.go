package main

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

// Output results to the console
func Output(results []Result) {
	if len(results) == 0 {
		log.Fatal("no results...")
	}

	var headers []string
	for key := range results[0].TestCase.Arguments {
		headers = append(headers, key)
	}

	table := tablewriter.NewWriter(os.Stdout)

	for _, result := range results {
		var row []string

		switch result.Status {
		case Success:
			row = append(row, "PASSED")
		case Failed:
			row = append(row, "FAILED")
		case Skipped:
			row = append(row, "SKIPPED")
		}

		row = append(row, result.TestCase.Client)
		row = append(row, result.Response)

		for _, header := range headers {
			row = append(row, result.TestCase.Arguments[header])
		}

		table.Append(row)
	}

	headers = append([]string{"status", "client", "response"}, headers...)
	table.SetHeader(headers)
	table.SetBorder(false)

	fmt.Println()
	table.Render() // Send output
	fmt.Println()
}
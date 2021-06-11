package models

import (
	"zendesk.com/interview/mrt-backend/core"
)

var lineMap map[string]int = map[string]int{}

// Extract all the lines, eg:["CC": 0, "EW: 0"]
func LineMap(path string) map[string]int {

	records := core.ReadCSV(path)
	prev := ""
	for _, line := range records {
		if prev != line[0][:2] {
			lineMap[prev] = 0
		}
		prev = line[0][:2]
	}
	return lineMap
}

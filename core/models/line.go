package models

import (
	"zendesk.com/interview/mrt-backend/core"
)

var stnLine map[string]int = map[string]int{}

func LineMapping(path string) map[string]int {

	records := core.ReadCSV(path)
	prev := ""
	for _, line := range records {
		if prev != line[0][:2] {
			stnLine[prev] = 0
		}
		prev = line[0][:2]
	}
	return stnLine
}

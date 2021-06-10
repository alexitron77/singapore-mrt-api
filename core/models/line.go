package models

import (
	"zendesk.com/interview/mrt-backend/core"
)

var stnLine map[string]int = map[string]int{}

func LineMapping() map[string]int {

	records := core.ReadCSV("../zendesk/public/station_map.csv")
	prev := ""
	for _, line := range records {
		if prev != line[0][:2] {
			stnLine[prev] = 0
		}
		prev = line[0][:2]
	}
	return stnLine
}

package models

import (
	"zendesk.com/interview/mrt-backend/core"
)

var stnMap map[string]string = map[string]string{}

// Build a dictonary of station with station code and station name as key pair values, eg: ["CC1": "Dhoby Ghaut", "CC10": "MacPherson"]
func StnMap(path string) map[string]string {
	records := core.ReadCSV(path)
	for _, line := range records {
		stnMap[line[0]] = line[1]
	}
	return stnMap
}

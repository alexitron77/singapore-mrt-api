package models

import (
	"zendesk.com/interview/mrt-backend/core"
)

var stationMapping map[string]string = map[string]string{}

// Build a dictonary of station with station code and station name as key pair values, eg: ["CC1": "Dhoby Ghaut", "CC10": "MacPherson"]
func StnMapping(path string) map[string]string {
	records := core.ReadCSV(path)
	for _, line := range records {
		stationMapping[line[0]] = line[1]
	}
	return stationMapping
}

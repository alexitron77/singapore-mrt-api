package models

import "zendesk.com/interview/mrt-backend/core"

var stationMapping map[string]string = map[string]string{}

// Build a dictonary of station with station code and station name as key pair values
func StnMapping() map[string]string {
	records := core.ReadCSV("../zendesk/public/station_map.csv")
	for _, line := range records {
		stationMapping[line[0]] = line[1]
	}

	return stationMapping
}

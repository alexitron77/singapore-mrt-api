package utils

import "zendesk.com/interview/mrt-backend/conf"

var stationMapping map[string]string = map[string]string{}

func StnMapping() map[string]string {
	records := conf.ReadCSV("../zendesk/assets/station_map.csv")
	for _, line := range records {
		stationMapping[line[0]] = line[1]
	}

	return stationMapping
}

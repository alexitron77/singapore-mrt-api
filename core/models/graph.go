package models

import (
	"zendesk.com/interview/mrt-backend/core"
)

var graph map[string][]string = map[string][]string{}
var stnMapping map[string][]string = map[string][]string{}

// Build a adjacent list graph based on a csv file of Singapore MRT station
func BuildGraph(path string) map[string][]string {
	if len(graph) > 0 {
		return graph
	}

	records := core.ReadCSV(path)
	prevStn := ""

	for index, line := range records {
		stnMapping[line[1]] = append(stnMapping[line[1]], line[0])

		if index == 0 {
			prevStn = line[0]
			continue
		}

		if line[0][:2] != prevStn[:2] {
			prevStn = line[0]
			continue
		}
		currStn := line[0]
		addEdge(prevStn, currStn)
		addEdge(currStn, prevStn)
		prevStn = currStn
	}
	addInterchange()
	return graph
}

// Add direct connected stations to the current station (left and right stations)
func addEdge(stn string, prevStn string) {
	graph[stn] = append(graph[stn], prevStn)
}

// Add different line stations for interchange stations
func addInterchange() {
	for _, stn := range stnMapping {
		for j, connectStn := range stn {
			temp := make([]string, len(stn))
			copy(temp, stn)
			slice := append(temp[:j], temp[j+1:]...)
			graph[connectStn] = append(graph[connectStn], slice...)
		}

	}
}

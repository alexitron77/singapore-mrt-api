package conf

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

var graph map[string][]string = map[string][]string{}

func BuildGraph() map[string][]string {
	if len(graph) > 0 {
		return graph
	}

	file, _ := filepath.Abs("../zendesk/assets/station_map.csv")
	csvFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully open file")
		defer csvFile.Close()
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	prevStn := ""

	for index, line := range csvLines {
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

	return graph
}

func addEdge(stn string, prevStn string) {
	graph[stn] = append(graph[stn], prevStn)
}

package utils

import (
	"fmt"

	"zendesk.com/interview/mrt-backend/core/models"
)

// Display itinerary in a human readable format
func PrintPaths(src string, dest string, paths [][]string, computedTravelTime []int) {
	path := "../zendesk/public/station_map.csv"
	stnMap := models.StnMap(path)
	fmt.Printf("Travel from %s to %s\n\n", stnMap[src], stnMap[dest])

	if len(paths) == 0 {
		fmt.Printf("No route found\n")
	}

	for i, path := range paths {
		if i == 0 {
			fmt.Printf("Best route: %s\n", path)
		} else {
			fmt.Printf("Alternative route: %s\n", path)
		}
		if len(computedTravelTime) > 0 {
			if computedTravelTime[i] == -1 {
				break
			}
			fmt.Printf("Total Travel time: %d\n\n", computedTravelTime[i])
		} else {
			fmt.Printf("Total Station travelled: %d\n\n", len(path))
		}

		prev := path[0]

		for _, curr := range path[1:] {
			if prev[:2] != curr[:2] {
				fmt.Printf("Change from %s line to %s line\n", prev[:2], curr[:2])
				prev = curr
				continue
			}
			fmt.Printf("Travel on %s line from %s to %s\n", prev[:2], stnMap[prev], stnMap[curr])
			prev = curr
		}
		fmt.Print("\n----\n\n")
	}
}

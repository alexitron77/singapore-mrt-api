package utils

import "fmt"

func ItineraryInfo(src string, dest string, paths [][]string) {
	stnMapping := StnMapping()
	fmt.Printf("Travel from %s to %s\n\n", stnMapping[src], stnMapping[dest])

	for i, path := range paths {
		if i == 0 {
			fmt.Printf("Best route: %s\n\n", path)
		} else {
			fmt.Printf("Alternative route: %s\n\n", path)
		}
		fmt.Printf("Total Station travelled: %d\n\n", len(path))

		prev := path[0]

		for _, curr := range path[1:] {
			if prev[:2] != curr[:2] {
				fmt.Printf("Change from %s line to %s line\n", prev[:2], curr[:2])
				prev = curr
				continue
			}
			fmt.Printf("Travel on %s line from %s to %s\n", prev[:2], stnMapping[prev], stnMapping[curr])
			prev = curr
		}
		fmt.Println()
	}
}

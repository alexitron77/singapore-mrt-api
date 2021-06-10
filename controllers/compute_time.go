package controllers

import "fmt"

func ComputeTime(paths [][]string, lineStnTime map[string]int, lineChangeTime int) {
	for _, path := range paths {
		totalTime := 0
		prev := path[0][:2]
		// fmt.Print(path)
		for _, stn := range path {
			if lineStnTime[stn[:2]] == 0 {
				break
			} else {
				if prev != stn[:2] {
					fmt.Print("station", stn)
					totalTime += lineChangeTime
					prev = stn[:2]
					continue
				}
				totalTime += lineStnTime[stn[:2]]
				prev = stn[:2]
			}
		}
		fmt.Print(totalTime)
	}
}

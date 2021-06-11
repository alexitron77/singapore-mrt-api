package controllers

// Calculate travel time based on shift timings
func ComputeTravelTime(paths [][]string, lineStnTime map[string]int, lineChangeTime int) []int {
	computedTravelTime := []int{}

	for _, path := range paths {
		totalTravelTime := 0
		prev := path[0][:2]
		for _, stn := range path {
			if lineStnTime[stn[:2]] == 0 {
				computedTravelTime = append(computedTravelTime, -1)
				break
			} else {
				if prev != stn[:2] {
					totalTravelTime += lineChangeTime
					prev = stn[:2]
					continue
				}
				totalTravelTime += lineStnTime[stn[:2]]
				prev = stn[:2]
			}
		}
		computedTravelTime = append(computedTravelTime, totalTravelTime)
	}
	return computedTravelTime
}

package controllers

// Calculate travel time based on shift timings
func ComputeTravelTime(paths [][]string, lineTimeConf map[string]int, lineChangeTime int) []int {
	computedTravelTime := []int{}

	for _, path := range paths {
		totalTravelTime := 0
		prev := path[0][:2]
		for _, stn := range path {
			// If the line is not operating (night-shift), we set the computedTravelTime to -1
			if lineTimeConf[stn[:2]] == 0 {
				computedTravelTime = append(computedTravelTime, -1)
				break
			} else {
				// From EW14 to CC9, compare EW and CC => A change of line is detected.
				if prev != stn[:2] {
					totalTravelTime += lineChangeTime
					prev = stn[:2]
					continue
				}
				// The mapping table (lineStnTime) is used to retrieve the travel time between stations on the line.
				totalTravelTime += lineTimeConf[stn[:2]]
				prev = stn[:2]
			}
		}
		computedTravelTime = append(computedTravelTime, totalTravelTime)
	}
	return computedTravelTime
}

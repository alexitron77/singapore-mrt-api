package controllers

import (
	"time"

	models "zendesk.com/interview/mrt-backend/core/models"
)

// Get shifts based on startTime
func GetTravelShift(paths [][]string, startTime string) (map[string]int, int, error) {
	start, err := time.Parse("2006-01-02T15:04", startTime)

	if err != nil {
		return nil, 0, err
	}

	weekday := start.Weekday().String()
	hour := start.Hour()

	// Compute the shift configuration based on weekday and hour
	lineTimeConf, lineChangeTime := computeShift(weekday, hour)

	return lineTimeConf, lineChangeTime, nil
}

func computeShift(weekday string, hour int) (map[string]int, int) {
	lineTimeConf := models.LineMap(path)
	lineChangeTime := 0

	// Peak hour
	if (hour >= 6 && hour < 9) && (weekday != "Saturday" && weekday != "Sunday") || ((weekday == "Saturday" || weekday == "Sunday") && (hour >= 18 && hour < 21)) {
		lineTimeConf["CC"] = 10
		lineTimeConf["CE"] = 10
		lineTimeConf["CG"] = 10
		lineTimeConf["DT"] = 10
		lineTimeConf["EW"] = 10
		lineTimeConf["NE"] = 12
		lineTimeConf["NS"] = 12
		lineTimeConf["TE"] = 10
		lineChangeTime = 15
	} else if hour >= 22 && hour < 6 { // Night hours
		lineTimeConf["CC"] = 10
		lineTimeConf["CE"] = 0
		lineTimeConf["CG"] = 0
		lineTimeConf["DT"] = 0
		lineTimeConf["EW"] = 10
		lineTimeConf["NE"] = 10
		lineTimeConf["NS"] = 10
		lineTimeConf["TE"] = 8
		lineChangeTime = 10
	} else {
		lineTimeConf["CC"] = 10
		lineTimeConf["CE"] = 10
		lineTimeConf["CG"] = 10
		lineTimeConf["DT"] = 8
		lineTimeConf["EW"] = 10
		lineTimeConf["NE"] = 10
		lineTimeConf["NS"] = 10
		lineTimeConf["TE"] = 8
		lineChangeTime = 10
	}

	return lineTimeConf, lineChangeTime

}

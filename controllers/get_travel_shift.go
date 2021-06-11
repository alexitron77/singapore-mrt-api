package controllers

import (
	"time"

	models "zendesk.com/interview/mrt-backend/core/models"
)

// Compute shifts based on startTime
func GetTravelShift(paths [][]string, startTime string) (map[string]int, int, error) {
	start, err := time.Parse("2006-01-02T15:04", startTime)

	if err != nil {
		return nil, 0, err
	}

	weekday := start.Weekday().String()
	hour := start.Hour()

	lineStationTime, lineChangeTime := computeShift(weekday, hour)

	return lineStationTime, lineChangeTime, nil
}

func computeShift(weekday string, hour int) (map[string]int, int) {
	lineMapping := models.LineMapping(path)
	lineChangeTime := 0

	// Peak hour
	if (hour >= 6 && hour < 9) && (weekday != "Saturday" && weekday != "Sunday") || ((weekday == "Saturday" || weekday == "Sunday") && (hour >= 18 && hour < 21)) {
		lineMapping["CC"] = 10
		lineMapping["CE"] = 10
		lineMapping["CG"] = 10
		lineMapping["DT"] = 10
		lineMapping["EW"] = 10
		lineMapping["NE"] = 12
		lineMapping["NS"] = 12
		lineMapping["TE"] = 10
		lineChangeTime = 15
	} else if hour >= 22 && hour < 6 { // Night hours
		lineMapping["CC"] = 10
		lineMapping["CE"] = 0
		lineMapping["CG"] = 0
		lineMapping["DT"] = 0
		lineMapping["EW"] = 10
		lineMapping["NE"] = 10
		lineMapping["NS"] = 10
		lineMapping["TE"] = 8
		lineChangeTime = 10
	} else {
		lineMapping["CC"] = 10
		lineMapping["CE"] = 10
		lineMapping["CG"] = 10
		lineMapping["DT"] = 8
		lineMapping["EW"] = 10
		lineMapping["NE"] = 10
		lineMapping["NS"] = 10
		lineMapping["TE"] = 8
		lineChangeTime = 10
	}

	return lineMapping, lineChangeTime

}

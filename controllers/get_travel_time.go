package controllers

import (
	"fmt"
	"time"

	models "zendesk.com/interview/mrt-backend/core/models"
)

func GetTravelTime(paths [][]string, startTime string) (map[string]int, int) {
	start, err := time.Parse("2006-01-02T15:04", startTime)

	if err != nil {
		fmt.Print("The start date format in incorrect")
	}

	weekday := start.Weekday().String()
	hour := start.Hour()

	lineStationTime, lineChangeTime := getShift(weekday, hour)

	return lineStationTime, lineChangeTime
}

func getShift(weekday string, hour int) (map[string]int, int) {
	lineMapping := models.LineMapping()
	lineChangeTime := 0

	// Peak hour
	if (hour > 6 && hour < 9) && (weekday != "Saturday" && weekday != "Sunday") || ((weekday == "Saturday" || weekday == "Sunday") && (hour > 18 && hour < 21)) {
		lineMapping["CC"] = 10
		lineMapping["CE"] = 10
		lineMapping["CG"] = 10
		lineMapping["DT"] = 10
		lineMapping["EW"] = 10
		lineMapping["NE"] = 12
		lineMapping["NS"] = 12
		lineMapping["TE"] = 10

		lineChangeTime = 15
	}
	return lineMapping, lineChangeTime

}

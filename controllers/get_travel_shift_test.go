package controllers

import (
	"reflect"
	"testing"
)

func init() {
	path = "../public/station_map.csv"
}
func TestGetTravelShift(t *testing.T) {

	paths := [][]string{
		{"DT11", "DT12", "NE7", "NE8", "NE9"},
		{"DT11", "DT10", "DT9", "CC19", "CC17", "CC16", "CC15", "CC14", "CC13", "NE12", "NE11", "NE10", "NE9"},
	}
	startTime := "2018-01-20T19:35"
	expectedLineMapping := map[string]int{
		"":   0,
		"CC": 10,
		"CE": 10,
		"CG": 10,
		"DT": 10,
		"EW": 10,
		"NE": 12,
		"NS": 12,
		"TE": 10,
	}

	expectedLineChangeTime := 15

	lineTime, changeTime, _ := GetTravelShift(paths, startTime)

	if reflect.DeepEqual(expectedLineMapping, lineTime) == false {
		t.Errorf("Expected %v. Got %v", expectedLineMapping, lineTime)
	} else {
		t.Logf("Successfully pass")
	}

	if expectedLineChangeTime != changeTime {
		t.Errorf("Expected %v\n Got %v", expectedLineChangeTime, changeTime)
	} else {
		t.Log("Successfully pass")
	}
}

func TestComputeShift(t *testing.T) {
	weekday := "Saturday"
	hour := 19

	lineMapping, lineChangeTime := computeShift(weekday, hour)

	expectedLineMapping := map[string]int{
		"":   0,
		"CC": 10,
		"CE": 10,
		"CG": 10,
		"DT": 10,
		"EW": 10,
		"NE": 12,
		"NS": 12,
		"TE": 10,
	}

	expectedLineChangeTime := 15

	if reflect.DeepEqual(expectedLineMapping, lineMapping) == false {
		t.Errorf("Expected %v\n Got %v", expectedLineMapping, lineMapping)
	} else {
		t.Logf("Successfully pass")
	}

	if reflect.DeepEqual(expectedLineChangeTime, lineChangeTime) == false {
		t.Errorf("Expected %v\n Got %v", expectedLineChangeTime, lineChangeTime)
	} else {
		t.Logf("Successfully pass")
	}

}

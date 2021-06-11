package controllers

import (
	"reflect"
	"testing"
)

func init() {
	path = "../public/station_map.csv"
}
func TestComputePath(t *testing.T) {

	expected := [][]string{
		{"DT11", "DT12", "NE7", "NE8", "NE9"},
		{"DT11", "DT10", "DT9", "CC19", "CC17", "CC16", "CC15", "CC14", "CC13", "NE12", "NE11", "NE10", "NE9"},
	}

	res := ComputePath("DT11", "NE9")

	if reflect.DeepEqual(expected, res) == false {
		t.Errorf("Expected %v\n Got %v", expected, res)
	} else {
		t.Logf("Successfully pass")
	}
}

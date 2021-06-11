package controllers

import (
	"reflect"
	"testing"
)

func TestComputeTime(t *testing.T) {

	paths := [][]string{
		{"DT11", "DT12", "NE7", "NE8", "NE9"},
		{"DT11", "DT10", "DT9", "CC19", "CC17", "CC16", "CC15", "CC14", "CC13", "NE12", "NE11", "NE10", "NE9"},
	}
	lineStnTime := map[string]int{"CC": 10, "CE": 10, "CG": 10, "DT": 10, "EW": 10, "NE": 12, "NS": 12, "TE": 10}
	lineChangeTime := 15
	expected := []int{59, 146}

	res := ComputeTravelTime(paths, lineStnTime, lineChangeTime)

	if reflect.DeepEqual(expected, res) == false {
		t.Errorf("Expected %v\n Got %v", expected, res)
	} else {
		t.Logf("Successfully pass")
	}
}

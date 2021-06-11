package models

import (
	"reflect"
	"testing"
)

func TestLineMap(t *testing.T) {
	expected := map[string]int{
		"":   0,
		"CC": 0,
		"CE": 0,
		"CG": 0,
		"DT": 0,
		"EW": 0,
		"NE": 0,
		"NS": 0,
	}

	path := "../../public/station_map.csv"

	res := LineMap(path)

	if reflect.DeepEqual(expected, res) == false {
		t.Errorf("Expected %v\n Got %v", expected, res)
	} else {
		t.Logf("Successfully pass")
	}
}

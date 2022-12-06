package lib

import "testing"

func max(arr ...int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if m < arr[i] {
			m = arr[i]
		}
	}
	return m
}

func min(arr ...int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if m > arr[i] {
			m = arr[i]
		}
	}
	return m
}

func assertEqual(t *testing.T, a, b interface{}, in interface{}) {
	if a != b {
		t.Log("in:", in)
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

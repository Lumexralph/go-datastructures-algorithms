package main

import (
	"fmt"
	"testing"
)

func TestPowerSeries(t *testing.T) {
	gotValueA, gotValueB, _ := powerSeries(5)
	expectedvalueA, expectedValueB := 25, 125

	if gotValueA != expectedvalueA {
		t.Errorf("powerSeries(5) = %d; want %d", gotValueA, expectedvalueA)
	}

	if gotValueB != expectedValueB {
		t.Errorf("powerSeries(5) = %d; want %d", gotValueB, expectedValueB)
	}
}

// create a table-driven version of the test
func TestPowerSeriesTwo(t *testing.T) {
	testCases := []struct {
		val            int
		expectedvalueA int
		expectedValueB int
	}{
		{2, 4, 8},
		{3, 9, 27},
		{1, 1, 0},
		{4, 16, 63},
	}

	// run the multiple testcases
	for _, tc := range testCases {

		t.Run(fmt.Sprintf("PowerSeries of %d", tc.val), func(t *testing.T) {

			gotValueA, gotValueB, _ := powerSeries(tc.val)

			if gotValueA != tc.expectedvalueA {
				t.Errorf("powerSeries(%d) = %d; want %d", tc.val, gotValueA, tc.expectedvalueA)
			}

			if gotValueB != tc.expectedValueB {
				t.Errorf("powerSeries(%d) = %d; want %d", tc.val, gotValueB, tc.expectedValueB)
			}

		})

	}
}

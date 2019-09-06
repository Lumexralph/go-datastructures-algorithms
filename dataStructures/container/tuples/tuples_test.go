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
		{1, 1, 1},
		{4, 16, 64},
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


// To know the performance of the powerSeries function
func BenchmarkPowerSeries(b *testing.B) {

	benchmarks := []struct{
		name string
		val int
		expectedvalueA int
		expectedValueB int
	} {
		{"Powerseries of 5", 5, 25, 125},
		{"Powerseries of 9", 9, 81, 729},
		{"Powerseries of 2", 2, 4, 8},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				powerSeries(bm.val)
			}
		})
	}
}
package main

import "testing"

func TestSafeRegionSize(t *testing.T) {
	tt := []struct {
		arg1 []xy
		arg2 int
		want int
	}{
		{[]xy{
			{1, 1},
			{1, 6},
			{8, 3},
			{3, 4},
			{5, 5},
			{8, 9},
		}, 32, 16},
	}
	for _, tc := range tt {
		t.Run("1", func(t *testing.T) {
			if got := safeRegionSize(tc.arg1, tc.arg2); got != tc.want {
				t.Errorf("safeRegionSize() = %d, want %d", got, tc.want)
			}
		})
	}
}

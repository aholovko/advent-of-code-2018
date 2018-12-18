package main

import "testing"

func TestLargestArea(t *testing.T) {
	tt := []struct {
		arg  []xy
		want int
	}{
		{[]xy{
			{1, 1},
			{1, 6},
			{8, 3},
			{3, 4},
			{5, 5},
			{8, 9},
		}, 17},
	}
	for _, tc := range tt {
		t.Run("1", func(t *testing.T) {
			if got := largestArea(tc.arg); got != tc.want {
				t.Errorf("largestArea() = %d, want %d", got, tc.want)
			}
		})
	}
}

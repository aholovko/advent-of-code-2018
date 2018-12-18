package main

import "testing"

func TestShortestPolymer(t *testing.T) {
	tt := []struct {
		arg  string
		want int
	}{
		{"aA", 0},
		{"abBA", 0},
		{"aabAAB", 0},
		{"dabAcCaCBAcCcaDA", 4},
	}
	for _, tc := range tt {
		t.Run("1", func(t *testing.T) {
			if got := shortestPolymer(tc.arg); got != tc.want {
				t.Errorf("shortestPolymer(%q) = %d, want %d", tc.arg, got, tc.want)
			}
		})
	}
}

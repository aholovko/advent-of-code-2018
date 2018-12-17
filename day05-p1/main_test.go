package main

import "testing"

func TestReactPolymer(t *testing.T) {
	tt := []struct {
		arg  string
		want int
	}{
		{"aA", 0},
		{"abBA", 0},
		{"aabAAB", 6},
		{"dabAcCaCBAcCcaDA", 10},
	}
	for _, tc := range tt {
		t.Run("1", func(t *testing.T) {
			if got := reactPolymer(tc.arg); got != tc.want {
				t.Errorf("reactPolymer(%q) = %d, want %d", tc.arg, got, tc.want)
			}
		})
	}
}

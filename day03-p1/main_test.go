package main

import "testing"

func TestOverlappedClaims(t *testing.T) {
	tt := []struct {
		arg  []claim
		want int
	}{
		{
			arg: []claim{
				{1, 3, 4, 4},
				{3, 1, 4, 4},
				{5, 5, 2, 2},
			},
			want: 4,
		},
	}
	for _, tc := range tt {
		t.Run("1", func(t *testing.T) {
			if got := overlappedClaims(tc.arg); got != tc.want {
				t.Errorf("overlappedClaims(%v) = %d, want %d", tc.arg, got, tc.want)
			}
		})
	}
}

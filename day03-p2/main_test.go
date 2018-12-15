package main

import "testing"

func TestFindNonOverlappingClaim(t *testing.T) {
	tt := []struct {
		arg  []claim
		want int
	}{
		{
			arg: []claim{
				{1, 1, 3, 4, 4},
				{2, 3, 1, 4, 4},
				{3, 5, 5, 2, 2},
			},
			want: 3,
		},
	}
	for _, tc := range tt {
		t.Run("1", func(t *testing.T) {
			fabric := fabric{}
			for _, c := range tc.arg {
				fabric.addClaim(c)
			}
			if got := fabric.findNonOverlappingClaim(); got != tc.want {
				t.Errorf("findNonOverlappingClaim(%v) = %d, want %d", tc.arg, got, tc.want)
			}
		})
	}
}

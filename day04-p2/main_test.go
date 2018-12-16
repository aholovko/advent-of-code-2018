package main

import (
	"testing"
)

func TestFindNonOverlappingClaim(t *testing.T) {
	tt := []struct {
		arg          journal
		wantGuard    int
		wantAsleepAt int
	}{
		{
			arg: journal{
				map[int][]shift{
					10: []shift{
						shift{
							date: "11-01",
							asleep: []interval{
								interval{start: 5, end: 25},
								interval{start: 30, end: 55},
							},
						},
						shift{
							date: "11-03",
							asleep: []interval{
								interval{start: 24, end: 29},
							},
						},
					},
					99: []shift{
						shift{
							date: "11-02",
							asleep: []interval{
								interval{start: 40, end: 50},
							},
						},
						shift{
							date: "11-04",
							asleep: []interval{
								interval{start: 36, end: 46},
							},
						},
						shift{
							date: "11-05",
							asleep: []interval{
								interval{start: 45, end: 55},
							},
						},
					},
				},
			},
			wantGuard:    99,
			wantAsleepAt: 45,
		},
	}
	for _, tc := range tt {
		t.Run("1", func(t *testing.T) {
			journal := tc.arg
			gotGuard, gotAsleepAt := journal.findMostAsleepGuardOnTheSameMin()
			if gotGuard != tc.wantGuard {
				t.Errorf("findMostAsleepGuardOnTheSameMin(), guard = %d, want %d", gotGuard, tc.wantGuard)
			}
			if gotAsleepAt != tc.wantAsleepAt {
				t.Errorf("findMostAsleepGuardOnTheSameMin(), asleepAt = %d, want %d", gotAsleepAt, tc.wantAsleepAt)
			}
		})
	}
}

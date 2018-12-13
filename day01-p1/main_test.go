package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFreqCalibration(t *testing.T) {
	tests := []struct {
		arg  []int
		want int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1, -1}, 0},
		{[]int{1, -1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(name(tt.arg), func(t *testing.T) {
			if got := freqCalibration(tt.arg); got != tt.want {
				t.Errorf("freqCalibration(%v) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

func name(nums []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), "_"), "[]")
}

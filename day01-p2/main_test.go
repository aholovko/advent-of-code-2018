package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFreqFind(t *testing.T) {
	tests := []struct {
		arg  []int
		want int
	}{
		{[]int{1, -1}, 0},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	}
	for _, tt := range tests {
		t.Run(name(tt.arg), func(t *testing.T) {
			if got := freqFind(tt.arg); got != tt.want {
				t.Errorf("freqFind(%v) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

func name(nums []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), "_"), "[]")
}

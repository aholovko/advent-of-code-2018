package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestChecksum(t *testing.T) {
	tests := []struct {
		arg  []string
		want int
	}{
		{[]string{}, 0},
		{[]string{"a"}, 0},
		{[]string{"aa"}, 0},
		{[]string{"aabbb"}, 1},
		{[]string{"aabbb", "aa", "bbb"}, 4},
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", strings.Join(tt.arg, "_")), func(t *testing.T) {
			if got := checksum(tt.arg); got != tt.want {
				t.Errorf("checksum(%v) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

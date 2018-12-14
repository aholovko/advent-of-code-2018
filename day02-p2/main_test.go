package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCommon(t *testing.T) {
	tt := []struct {
		arg  []string
		want string
	}{
		{[]string{"abcde", "abcde"}, ""},
		{[]string{"abcde", "fghij"}, ""},
		{[]string{"fghij", "fguij"}, "fgij"},
		{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s", strings.Join(tc.arg, "_")), func(t *testing.T) {
			if got := common(tc.arg); got != tc.want {
				t.Errorf("common(%v) = %q, want %q", tc.arg, got, tc.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestStepsOrder(t *testing.T) {
	tt := []struct {
		arg  map[rune][]rune
		want []rune
	}{
		// {
		// 	map[rune][]rune{},
		// 	[]rune{},
		// },
		// {
		// 	map[rune][]rune{
		// 		'A': []rune{'B'},
		// 	},
		// 	[]rune{'A', 'B'},
		// },
		// {
		// 	map[rune][]rune{
		// 		'A': []rune{'B'},
		// 		'B': []rune{'C'},
		// 	},
		// 	[]rune{'A', 'B', 'C'},
		// },
		// {
		// 	map[rune][]rune{
		// 		'A': []rune{'C', 'D'},
		// 		'C': []rune{'B'},
		// 		'D': []rune{'B'},
		// 	},
		// 	[]rune{'A', 'C', 'D', 'B'},
		// },
		// {
		// 	map[rune][]rune{
		// 		'A': []rune{'C', 'D'},
		// 		'C': []rune{'B'},
		// 	},
		// 	[]rune{'A', 'C', 'B', 'D'},
		// },
		// {
		// 	map[rune][]rune{
		// 		'A': []rune{'C', 'B'},
		// 		'B': []rune{'D', 'E'},
		// 	},
		// 	[]rune{'A', 'B', 'C', 'D', 'E'},
		// },
		// {
		// 	map[rune][]rune{
		// 		'C': []rune{'A', 'F'},
		// 		'A': []rune{'B', 'D'},
		// 		'B': []rune{'E'},
		// 		'D': []rune{'E'},
		// 		'F': []rune{'E'},
		// 	},
		// 	[]rune{'C', 'A', 'B', 'D', 'F', 'E'},
		// },
		// {
		// 	map[rune][]rune{
		// 		'C': []rune{'B', 'H', 'F'},
		// 		'B': []rune{'D'},
		// 		'H': []rune{'D', 'E'},
		// 		'F': []rune{'A'},
		// 		'E': []rune{'G'},
		// 	},
		// 	[]rune{'C', 'B', 'D', 'F', 'A', 'H', 'E', 'G'},
		// },
		{
			map[rune][]rune{
				'E': []rune{'B', 'F'},
				'B': []rune{'A', 'C'},
				'A': []rune{'D'},
				'F': []rune{'D', 'H'},
				'I': []rune{'H', 'K'},
				'K': []rune{'H'},
				'H': []rune{'J'},
				'D': []rune{'G'},
				'C': []rune{'G'},
				'J': []rune{'G'},
			},
			[]rune{'C', 'B', 'D', 'F', 'A', 'H', 'E', 'G'},
		},
	}
	for i, tc := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := stepsOrder(tc.arg); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("stepsOrder() = %q, want %q", string(got), string(tc.want))
			}
		})
	}
}

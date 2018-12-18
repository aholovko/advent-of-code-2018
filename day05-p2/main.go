package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	polymer := string(bytes)
	fmt.Printf("shortestPolymer() = %d\n", shortestPolymer(polymer))
}

func shortestPolymer(polymer string) int {
	uu := map[rune]int{} // unit -> shortest polymer after removing unit
	for _, u := range polymer {
		_, ok := uu[unicode.ToUpper(u)]
		if ok {
			continue
		}
		reactor := make([]rune, 0, len(polymer))
		r := rune(0)
		for _, c := range polymer {
			if c == unicode.ToUpper(u) || c == unicode.ToLower(u) {
				continue
			}
			if react(r, c) && len(reactor) > 0 {
				reactor = reactor[:len(reactor)-1]
				if len(reactor) > 0 {
					r = reactor[len(reactor)-1]
				}
				continue
			}
			reactor = append(reactor, c)
			r = c
		}
		uu[unicode.ToUpper(u)] = len(reactor)
	}

	min := len(polymer)
	for _, length := range uu {
		if length < min {
			min = length
		}
	}
	return min
}

func react(a, b rune) bool {
	return a != b && unicode.ToUpper(a) == unicode.ToUpper(b)
}

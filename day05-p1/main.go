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
	fmt.Printf("reactPolymer() = %d\n", reactPolymer(polymer))
}

func reactPolymer(polymer string) int {
	reactor := make([]rune, 0, len(polymer))
	r := rune(0)
	for _, c := range polymer {
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
	return len(reactor)
}

func react(a, b rune) bool {
	return a != b && unicode.ToUpper(a) == unicode.ToUpper(b)
}

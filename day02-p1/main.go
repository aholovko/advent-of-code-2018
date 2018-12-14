package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	words := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
	fmt.Printf("%d\n", checksum(words))
}

func checksum(words []string) int {
	var twos, threes int
	for _, w := range words {
		counts := map[rune]int{}
		for _, r := range w {
			counts[r]++
		}
		var hasTwos, hasThrees bool
		for _, c := range counts {
			switch c {
			case 2:
				hasTwos = true
			case 3:
				hasThrees = true
			}
		}
		if hasTwos {
			twos++
		}
		if hasThrees {
			threes++
		}
	}
	return twos * threes
}

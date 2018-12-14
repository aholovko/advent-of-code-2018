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

	strs := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		strs = append(strs, s.Text())
	}
	fmt.Printf("%d\n", checksum(strs))
}

func checksum(strs []string) int {
	if len(strs) == 0 {
		return 0
	}
	var twos, threes int
	for _, str := range strs {
		n2, n3 := stats(str)
		twos += n2
		threes += n3
	}
	return twos * threes
}

func stats(str string) (twos, threes int) {
	stats := map[rune]int{}
	for _, c := range str {
		stats[c]++
	}
	for _, n := range stats {
		switch n {
		case 2:
			twos = 1
		case 3:
			threes = 1
		}
	}
	return twos, threes
}

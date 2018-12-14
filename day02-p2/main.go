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
	fmt.Printf("%s\n", common(words))
}

// common finds common letters between the first two correct box IDs (words).
// Correct box IDs differ by exactly one character.
func common(words []string) string {
	for i, a := range words {
		for _, b := range words[i+1:] {
			idx := -1
			for j := 0; j < len(a); j++ {
				if a[j] == b[j] {
					continue
				}
				if idx != -1 {
					idx = -1
					break
				}
				idx = j
			}
			if idx != -1 {
				return a[:idx] + a[idx+1:]
			}
		}
	}
	return ""
}

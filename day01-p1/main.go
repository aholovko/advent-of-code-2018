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

	nums := []int{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		nums = append(nums, n)
	}

	fmt.Printf("%d\n", freqCalibration(nums))
}

func freqCalibration(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

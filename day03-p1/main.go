package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type claim struct {
	left   int
	top    int
	width  int
	height int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	claims := []claim{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n, left, top, width, height int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &n, &left, &top, &width, &height) // #1 @ 432,394: 29x14
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		claims = append(claims, claim{left, top, width, height})
	}
	fmt.Printf("%d\n", overlappedClaims(claims))
}

func overlappedClaims(claims []claim) int {
	fabric := [1010][1010]int{}
	for _, claim := range claims {
		for row := claim.top; row < claim.top+claim.height; row++ {
			for col := claim.left; col < claim.left+claim.width; col++ {
				fabric[row][col]++
			}
		}
	}
	overlapped := 0
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric[i]); j++ {
			if fabric[i][j] > 1 {
				overlapped++
			}
		}
	}
	return overlapped
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type xy struct{ x, y int }

func (p xy) distance(x, y int) int {
	return abs(p.x-x) + abs(p.y-y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	xys := []xy{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y int
		_, err := fmt.Sscanf(s.Text(), "%d, %d", &x, &y)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		xys = append(xys, xy{x, y})
	}
	fmt.Printf("largest area: %d\n", largestArea(xys))
}

func largestArea(xys []xy) int {
	maxX, maxY := 0, 0
	for _, p := range xys {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	grid := make([][]int, maxY+1)
	for i := range grid {
		grid[i] = make([]int, maxX+1)
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			min := int(^uint(0) >> 1) // max int
			idx := -1
			for i, p := range xys {
				d := p.distance(x, y) // Manhattan distance (l1 norm)
				if d <= min {
					if d == min {
						idx = -1 // more than 1 point with the same min distance
						continue
					}
					min = d
					idx = i
				}
			}
			grid[y][x] = idx
		}
	}

	coords := map[xy]int{} // point -> area
	for _, c := range xys {
		coords[c] = 0
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			n := grid[y][x]
			if n == -1 {
				continue
			}
			if x == 0 || y == 0 || x == maxX || y == maxY {
				p := xys[n]
				if _, ok := coords[p]; ok {
					delete(coords, p)
				}
			}
			p := xys[n]
			if _, ok := coords[p]; ok {
				coords[p]++
			}
		}
	}

	maxArea := 0
	for _, s := range coords {
		if s > maxArea {
			maxArea = s
		}
	}

	return maxArea
}

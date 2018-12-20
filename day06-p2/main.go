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
	fmt.Printf("Safe region size: %d\n", safeRegionSize(xys, 10000))
}

func safeRegionSize(xys []xy, maxDistance int) int {
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
			sum := 0
			for _, p := range xys {
				sum += p.distance(x, y)
			}
			grid[y][x] = sum
		}
	}

	size := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[y][x] < maxDistance {
				size++
			}
		}
	}

	return size
}

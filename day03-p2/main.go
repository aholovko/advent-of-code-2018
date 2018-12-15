package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type xy struct{ x, y int }

type fabric struct {
	m map[xy][]int
}

type claim struct {
	id, x, y, w, h int
}

func (f *fabric) addClaim(c claim) {
	if f.m == nil {
		f.m = map[xy][]int{}
	}
	for x := c.x - 1; x < c.x+c.w-1; x++ {
		for y := c.y - 1; y < c.y+c.h-1; y++ {
			f.m[xy{x, y}] = append(f.m[xy{x, y}], c.id)
		}
	}
}

func (f *fabric) display() {
	var maxX, maxY int
	for p := range f.m {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			fmt.Print(len(f.m[xy{x, y}]))
		}
		fmt.Println()
	}
}

func (f *fabric) overlappingArea() int {
	area := 0
	for _, claims := range f.m {
		if len(claims) > 1 {
			area++
		}
	}
	return area
}

func (f *fabric) findNonOverlappingClaim() int {
	claims := map[int]bool{}
	for _, p := range f.m {
		if len(p) > 1 {
			for _, c := range p {
				claims[c] = false
			}
			continue
		}
		c := p[0]
		_, ok := claims[c]
		if !ok {
			claims[c] = true
		}
	}
	for c, ok := range claims {
		if ok {
			return c
		}
	}
	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fabric := fabric{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h) // #1 @ 432,394: 29x14
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		fabric.addClaim(claim{id, x, y, w, h})
	}
	fmt.Printf("%d\n", fabric.findNonOverlappingClaim())
}

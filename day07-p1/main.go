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

	steps := map[rune][]rune{} // step -> dependant steps
	s := bufio.NewScanner(f)
	for s.Scan() {
		var a, b string
		_, err := fmt.Sscanf(s.Text(), "Step %s must be finished before step %s can begin.", &a, &b)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		k := []rune(a)[0]
		if n, ok := steps[k]; ok {
			steps[k] = append(n, []rune(b)[0])
			continue
		}
		steps[k] = []rune{[]rune(b)[0]}
	}
	fmt.Printf("Steps: %s\n", string(stepsOrder(steps)))
}

func stepsOrder(steps map[rune][]rune) []rune {
	// Kahn's algorithm (https://en.m.wikipedia.org/wiki/Topological_sorting)
	// https://www.youtube.com/watch?v=tFpvX8T0-Pw
	l := []rune{} // list of ordered steps
	s := []rune{} // steps with no incoming dependencies
	n := rune(0)

	dep := map[rune]bool{} // step -> exist dependant step?
	for k := range steps {
		dep[k] = false
	}
	for _, v := range steps {
		for _, s := range v {
			dep[s] = true
		}
	}

	for k, v := range dep {
		if v {
			continue
		}
		s = add(s, k)
	}

	for len(s) > 0 {
		n, s = s[0], s[1:]
		l = append(l, n)
		for _, m := range steps[n] {
			steps[n] = remove(steps[n], m)
			fmt.Println(steps)
			hasIncoming := false
			for _, v := range steps {
				if contains(v, m) {
					hasIncoming = true
					break
				}
			}
			if !hasIncoming {
				s = add(s, m)
			}
		}
	}
	return l
}

func contains(a []rune, x rune) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func remove(a []rune, x rune) []rune {
	for i, v := range a {
		if v == x {
			a = append(a[:i], a[i+1:]...)
			return a
		}
	}
	return a
}

func add(a []rune, x rune) []rune {
	for i, v := range a {
		if v == x {
			return a
		}
		if v > x {
			a = append(a, 0)
			copy(a[i+1:], a[i:])
			a[i] = x
			return a
		}
	}
	a = append(a, x)
	return a
}

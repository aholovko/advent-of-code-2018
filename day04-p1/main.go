package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type journal struct {
	m map[int][]shift // guard -> shifts
}

type shift struct {
	date   string
	asleep []interval
}

type interval struct{ start, end int }

func (j *journal) findMostAsleepGuard() (guard int, asleepAt int) {
	asleepTotal := 0
	for id, shifts := range j.m {
		slept := 0
		for _, shift := range shifts {
			for _, t := range shift.asleep {
				slept += (t.end - t.start)
			}
		}
		if slept > asleepTotal {
			guard = id
			asleepTotal = slept
		}
	}

	minutes := [60]int{}
	for _, s := range j.m[guard] {
		for _, in := range s.asleep {
			for i := in.start; i < in.end; i++ {
				minutes[i]++
			}
		}
	}

	max := minutes[0]
	for i, v := range minutes {
		if v > max {
			asleepAt = i
			max = v
		}
	}

	return guard, asleepAt
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	journal := journal{
		m: map[int][]shift{},
	}

	entries := map[time.Time]string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		entry := s.Text() // [1518-11-01 00:00] Guard #10 begins shift
		idx := strings.Index(entry, "]")

		t, err := time.Parse("2006-01-02 15:04", entry[1:idx])
		if err != nil {
			log.Fatalf("could not parse time %s: %v", entry[1:idx], err)
		}

		entries[t] = strings.TrimSpace(entry[idx+1:])
	}

	var times []time.Time
	for t := range entries {
		times = append(times, t)
	}
	sort.Slice(times, func(i, j int) bool {
		return times[i].Before(times[j])
	})

	guard := 0
	start := 0
	date := ""
	intervals := []interval{}
	for _, t := range times {
		entry := entries[t]

		switch entry[:5] {
		case "Guard":
			var id int
			_, err := fmt.Sscanf(entry, "Guard #%d begins shift", &id)
			if err != nil {
				log.Fatalf("could not parse %s: %v", entry, err)
			}
			if guard != 0 {
				s := shift{
					date:   date,
					asleep: intervals,
				}
				ss, ok := journal.m[guard]
				if ok {
					journal.m[guard] = append(ss, s)
				} else {
					journal.m[guard] = []shift{s}
				}
				start = 0
				date = ""
				intervals = []interval{}
			}
			guard = id
		case "falls":
			start = t.Minute()
		case "wakes":
			date = fmt.Sprintf("%02d-%02d", t.Month(), t.Day())
			intervals = append(intervals, interval{
				start: start,
				end:   t.Minute(),
			})
		}
	}

	guard, asleep := journal.findMostAsleepGuard()
	fmt.Printf("Guard: %d, asleep @ %d\n", guard, asleep)
}

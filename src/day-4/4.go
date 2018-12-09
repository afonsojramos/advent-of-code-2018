package main

import (
	"fmt"
	"sort"
	"strings"
	"utils"
)

type guard struct {
	id     int
	shifts []shift
}

type shift struct {
	starthour, startmin int
	naps                []nap
}

type nap struct {
	asleep, awake int
}

func main() {
	// Part 1
	lines := utils.ReadLines("day-4/4.input")
	sort.Strings(lines)
	guards := make(map[int]guard)

	for i := 0; i < len(lines); {
		var month, day, hour, minute, id int
		utils.Sscanf(lines[i], "[1518-%d-%d %d:%d] Guard #%d begins shift", &month, &day, &hour, &minute, &id)

		var naps []nap
		for i++; i < len(lines) && !strings.Contains(lines[i], "begins"); i += 2 {
			var asleephour, asleepmin, wakehour, wakemin int
			utils.Sscanf(lines[i], "[1518-%d-%d %d:%d] falls asleep", &month, &day, &asleephour, &asleepmin)
			utils.Sscanf(lines[i+1], "[1518-%d-%d %d:%d] wakes up", &month, &day, &wakehour, &wakemin)
			naps = append(naps, nap{asleepmin, wakemin})
		}

		g := guards[id]
		g.id = id
		g.shifts = append(g.shifts, shift{hour, minute, naps})
		guards[id] = g
	}

	var max, maxID int
	for _, g := range guards {
		sleepTime := 0
		for _, shift := range g.shifts {
			for _, nap := range shift.naps {
				sleepTime += nap.awake - nap.asleep
			}
		}
		if sleepTime > max {
			max = sleepTime
			maxID = g.id
		}
	}

	minuteSleep := make(map[int]int)
	for _, shift := range guards[maxID].shifts {
		for _, nap := range shift.naps {
			for i := nap.asleep; i < nap.awake; i++ {
				minuteSleep[i]++
			}
		}
	}

	tempMax := 0
	maxMin := 0
	for minute, freq := range minuteSleep {
		if freq > tempMax {
			maxMin = minute
			tempMax = freq
		}
	}

	fmt.Println(maxMin, maxID, maxMin*maxID)
}

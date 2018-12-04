package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

const inputFileName = "input.txt"

type log struct {
	time   time.Time
	action string
}

type logs []log

func (l logs) Len() int {
	return len(l)
}

func (l logs) Less(i, j int) bool {
	return l[i].time.UnixNano() < l[j].time.UnixNano()
}

func (l logs) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var logs logs

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		var log log

		log.time, err = time.Parse("2006-01-02 15:04", line[1:17])
		if err != nil {
			panic(err)
		}

		log.action = line[19:]

		logs = append(logs, log)
	}

	sort.Sort(logs)

	var guardID int
	var slept time.Time
	guardSleepMinutes := make(map[int][60]int)

	for _, log := range logs {
		switch log.action[:5] {
		case "Guard":
			_, err := fmt.Sscanf(log.action, "Guard #%d begins shift", &guardID)
			if err != nil {
				panic(err)
			}
		case "falls":
			slept = log.time
		case "wakes":
			for i := slept; i.Unix() < log.time.Unix(); i = i.Add(time.Minute) {
				minutes := guardSleepMinutes[guardID]
				minutes[i.Minute()]++
				guardSleepMinutes[guardID] = minutes
			}
		}
	}

	var mostMinutesSlept int
	var sleepiestGuard int
	var sleepiestMinute int

	for guardID, sleepMinutes := range guardSleepMinutes {
		var minutesSlept int

		var mostSleptMinute int
		var mostSleptMinuteCount int
		for minute, count := range sleepMinutes {
			minutesSlept += count

			if count > mostSleptMinuteCount {
				mostSleptMinute = minute
				mostSleptMinuteCount = count
			}
		}

		if minutesSlept > mostMinutesSlept {
			mostMinutesSlept = minutesSlept
			sleepiestGuard = guardID
			sleepiestMinute = mostSleptMinute
		}
	}

	fmt.Println("sleepiest guard id multiplied by their sleepiest minute:", sleepiestGuard*sleepiestMinute)

	var mostSleptMinuteGuard int
	var mostSleptMinute int
	var mostSleptMinuteCount int

	for guardID, sleepMinutes := range guardSleepMinutes {
		for minute, count := range sleepMinutes {
			if count > mostSleptMinuteCount {
				mostSleptMinute = minute
				mostSleptMinuteCount = count
				mostSleptMinuteGuard = guardID
			}
		}
	}

	fmt.Println("guard who slept the most on a single minute id multiplied by that sleepiest minute:", mostSleptMinuteGuard*mostSleptMinute)
}

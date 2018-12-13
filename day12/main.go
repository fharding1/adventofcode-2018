package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	curGeneration := make(map[int]bool)

	scanner.Scan()
	for i, v := range strings.TrimPrefix(scanner.Text(), "initial state: ") {
		curGeneration[i] = v == '#'
	}
	scanner.Scan()

	rules := make(map[string]bool)
	for scanner.Scan() {
		rule := scanner.Text()
		rules[rule[:5]] = rule[9:10][0] == '#'
	}

	var minIdx, maxIdx = 0, len(curGeneration) - 1
	var stabilizedGrowthRate int64

	var generation int
	for generation = 0; generation < 100; generation++ {
		if generation == 20 {
			fmt.Println("sum of all indexes of plants that contain a plant after 20 generations:", sum(curGeneration))
		}

		nextGeneration := make(map[int]bool)
		stabilizedGrowthRate = 0

		start := minIdx - 2
		end := maxIdx + 2

		for plantIdx := start; plantIdx <= end; plantIdx++ {
			var test string
			for i := plantIdx - 2; i <= plantIdx+2; i++ {
				if curGeneration[i] {
					test += "#"
				} else {
					test += "."
				}
			}

			old := nextGeneration[plantIdx]

			res, ok := rules[test]
			if ok {
				nextGeneration[plantIdx] = res

				if plantIdx < minIdx {
					minIdx = plantIdx
				}

				if plantIdx > maxIdx {
					maxIdx = plantIdx
				}
			}

			if !old && nextGeneration[plantIdx] {
				stabilizedGrowthRate++
			}
		}

		for k, v := range nextGeneration {
			curGeneration[k] = v
		}
	}

	sum := sum(curGeneration) + int64(50000000000-generation)*stabilizedGrowthRate
	fmt.Println("sum of all indexes of plants that contain a plant after 50000000000 generations:", sum)
}

// sums the keys of true values in the map
func sum(m map[int]bool) int64 {
	var sum int64
	for i, v := range m {
		if v {
			sum += int64(i)
		}
	}
	return sum
}

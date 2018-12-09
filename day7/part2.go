package main

import (
	"bufio"
	"fmt"
	"os"
)

type worker struct {
	step      byte
	remaining int
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	dependencies := make(map[byte][]byte)
	for scanner.Scan() {
		line := scanner.Text()

		var step, dependency byte
		fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &dependency, &step)

		dependencies[step] = append(dependencies[step], dependency)

		if len(dependencies[dependency]) == 0 {
			dependencies[dependency] = make([]byte, 0)
		}
	}

	var workers [5]worker
	var sec int

	progress := make(map[byte]bool)

	for ; len(dependencies) > 0; sec++ {
		for i := range workers {
			workers[i].remaining--

			if workers[i].remaining <= 0 {
				progress[workers[i].step] = true
				delete(dependencies, workers[i].step)

				var choice byte
			DependencyLoop:
				for step, deps := range dependencies {
					for _, dep := range deps {
						if _, ok := dependencies[dep]; ok {
							continue DependencyLoop
						}
					}

					if _, ok := progress[step]; !ok && (step < choice || choice == 0) {
						choice = step
					}
				}

				if choice == 0 {
					continue
				}

				if _, ok := progress[choice]; !ok {
					workers[i].step = choice
					workers[i].remaining = int(choice) - 4
					progress[choice] = false
				}
			}
		}
	}

	fmt.Println("time to complete all the steps with 5 workers and the 60+ second step durations:", sec-1)
}

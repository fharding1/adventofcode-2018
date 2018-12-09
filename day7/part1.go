package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	dependencies := make(map[rune][]rune)
	for scanner.Scan() {
		line := scanner.Text()

		var step, dependency rune
		fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &dependency, &step)

		dependencies[step] = append(dependencies[step], dependency)

		if len(dependencies[dependency]) == 0 {
			dependencies[dependency] = make([]rune, 0)
		}
	}

	var ordered []rune

	for len(dependencies) > 0 {
		var choice rune

	DependencyLoop:
		for step, deps := range dependencies {
			for _, dep := range deps {
				if _, ok := dependencies[dep]; ok {
					continue DependencyLoop
				}
			}

			if step < choice || choice == 0 {
				choice = step
			}
		}

		ordered = append(ordered, choice)
		delete(dependencies, choice)
	}

	fmt.Println("order in which the instructions should be completed:", string(ordered))
}

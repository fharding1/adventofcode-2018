package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const inputFileName = "input.txt"

func main() {
	frequencies, err := parseFile(inputFileName)
	if err != nil {
		panic(err)
	}

	if len(frequencies) == 0 {
		return
	}

	var sum int
	history := map[int]struct{}{0: struct{}{}}

	for i := 0; true; i++ {
		sum += frequencies[i%len(frequencies)]

		if _, ok := history[sum]; ok {
			break
		}

		history[sum] = struct{}{}

		if i == len(frequencies)-1 {
			fmt.Println("first sum:", sum)
		}
	}

	fmt.Println("first duplicate sum:", sum)
}

func parseFile(name string) ([]int, error) {
	input, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(input), "\n")

	var numbers []int
	for _, line := range lines {
		if line == "" {
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, n)
	}

	return numbers, nil
}

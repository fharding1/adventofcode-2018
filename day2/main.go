package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const inputFileName = "input.txt"

func main() {
	lines, err := parseFile(inputFileName)
	if err != nil {
		panic(err)
	}

	counts := make(map[int]int)

	for _, line := range lines {
		letterCounts := make(map[rune]int)
		for _, ch := range line {
			if letterCounts[ch] == 0 {
				letterCounts[ch] = 1
			} else {
				letterCounts[ch]++
			}
		}

		counted := make(map[int]struct{})
		for _, v := range letterCounts {
			if _, ok := counted[v]; !ok {
				counted[v] = struct{}{}
				counts[v]++
			}
		}
	}

	fmt.Println("checksum:", counts[2]*counts[3])

A:
	for _, a := range lines {
	B:
		for _, b := range lines {
			off := -1

			for i := range b {
				if a[i] != b[i] {
					if off == -1 {
						off = i
					} else {
						continue B
					}
				}
			}

			if off != -1 {
				fmt.Println("common id:", a[:off]+a[off+1:])
				break A
			}
		}
	}
}

func parseFile(name string) ([]string, error) {
	input, err := ioutil.ReadFile(name)
	return strings.Split(string(input), "\n"), err
}

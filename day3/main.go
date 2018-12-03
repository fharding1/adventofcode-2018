package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const inputFileName = "input.txt"

type Claim struct {
	ID         int
	OffsetLeft int
	OffsetTop  int
	Width      int
	Height     int
}

func main() {
	claims, err := parseFile(inputFileName)
	if err != nil {
		panic(err)
	}

	var fabric [1000][1000][]int
	for _, c := range claims {
		for i := c.OffsetLeft; i < c.OffsetLeft+c.Width; i++ {
			for j := c.OffsetTop; j < c.OffsetTop+c.Height; j++ {
				fabric[i][j] = append(fabric[i][j], c.ID)
			}
		}
	}

	var count int
	for _, col := range fabric {
		for _, cell := range col {
			if len(cell) >= 2 {
				count++
			}
		}
	}

	fmt.Println("square inches of fabric within two or more claims:", count)

ClaimsLoop:
	for _, c := range claims {
		for i := c.OffsetLeft; i < c.OffsetLeft+c.Width; i++ {
			for j := c.OffsetTop; j < c.OffsetTop+c.Height; j++ {
				if len(fabric[i][j]) > 1 {
					continue ClaimsLoop
				}
			}
		}

		fmt.Println("non-overlapping claim id:", c.ID)
	}
}

func parseFile(name string) ([]Claim, error) {
	input, err := ioutil.ReadFile(name)

	lines := strings.Split(string(input), "\n")
	regex := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

	var claims []Claim
	for _, line := range lines {
		if line == "" {
			continue
		}

		matches := regex.FindAllStringSubmatch(line, -1)
		id, _ := strconv.Atoi(matches[0][1])
		ol, _ := strconv.Atoi(matches[0][2])
		ot, _ := strconv.Atoi(matches[0][3])
		w, _ := strconv.Atoi(matches[0][4])
		h, _ := strconv.Atoi(matches[0][5])
		c := Claim{id, ol, ot, w, h}
		claims = append(claims, c)
	}

	return claims, err
}

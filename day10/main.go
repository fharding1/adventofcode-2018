package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type star struct {
	x    int
	y    int
	velX int
	velY int
}

var re = regexp.MustCompile(`position=< *(-?\d+), *(-?\d+)> velocity=< *(-?\d+), *(-?\d+)>`)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	var sky []star

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)

		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[0][2])
		velX, _ := strconv.Atoi(matches[0][3])
		velY, _ := strconv.Atoi(matches[0][4])

		sky = append(sky, star{x, y, velX, velY})
	}

	var minBB boundingBox
	var minBBSky []star
	var seconds int

	for i := 0; i < 15000; i++ {
		bb := boundingBox{math.MaxInt64, math.MinInt64, math.MaxInt64, math.MinInt64}
		for _, v := range sky {
			if v.x < bb.minX {
				bb.minX = v.x
			}
			if v.x > bb.maxX {
				bb.maxX = v.x
			}
			if v.y < bb.minY {
				bb.minY = v.y
			}
			if v.y > bb.maxY {
				bb.maxY = v.y
			}
		}

		if bb.Size() < minBB.Size() || i == 0 {
			minBB = bb
			seconds = i
			minBBSky = make([]star, len(sky))
			copy(minBBSky, sky)
		}

		for i := range sky {
			sky[i].x += sky[i].velX
			sky[i].y += sky[i].velY
		}
	}

	display := make([][]bool, abs(minBB.maxY-minBB.minY)+1)
	for i := range display {
		display[i] = make([]bool, abs(minBB.maxX-minBB.minX)+1)
	}

	for _, star := range minBBSky {
		display[star.y-abs(minBB.minY)][star.x-abs(minBB.minX)] = true
	}

	fmt.Println("message that will eventually appear in the sky:")
	for y := range display {
		for _, ok := range display[y] {
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println("seconds they have to wait for that message to appear:", seconds)
}

type boundingBox struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (bb boundingBox) Size() int {
	return abs(bb.maxX-bb.minX) * abs(bb.maxY-bb.minY)
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}

	return a
}

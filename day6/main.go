// this is shit
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func manhattanDist(a, b, c, d float64) float64 {
	return float64(math.Abs(a-c) + math.Abs(b-d))
}

func main() {
	coords, err := parseFile("input.txt")
	if err != nil {
		panic(err)
	}

	var grid [500][500]int
	for i := range grid {
		for j := range grid[i] {
			distances := make(map[int]float64)

			for coordIdx, coord := range coords {
				dist := manhattanDist(float64(i), float64(j), float64(coord[0]), float64(coord[1]))
				distances[coordIdx] = dist
			}

			var smallest = math.MaxFloat64
			var smallestIdx int

			for coordIdx, dist := range distances {
				if dist < smallest {
					smallest = dist
					smallestIdx = coordIdx
				}
			}

			for coordIdx, dist := range distances {
				if dist == smallest && smallestIdx != coordIdx {
					smallestIdx = -1
				}
			}

			grid[i][j] = smallestIdx
		}
	}

	areas := make(map[int]int)
	infinite := make(map[int]bool)
	for i := range grid {
		for j := range grid[i] {
			coordIdx := grid[i][j]

			if coordIdx == -1 {
				continue
			}

			areas[grid[i][j]]++
			if i == 0 || i == len(grid)-1 || j == 0 || j == len(grid)-1 {
				infinite[grid[i][j]] = true
			}
		}
	}

	var biggestArea int

	for coordIdx, v := range areas {
		if v > biggestArea && !infinite[coordIdx] {
			biggestArea = v
		}
	}

	fmt.Println("size of the largest area that isn't infinite:", biggestArea)

	var size int
	for i := range grid {
		for j := range grid[i] {
			var totalDist float64
			for _, coord := range coords {
				dist := manhattanDist(float64(i), float64(j), float64(coord[0]), float64(coord[1]))
				totalDist += dist
			}

			if totalDist < 10000 {
				size++
			}
		}
	}

	fmt.Println("size of the region containing all locations which have a total distance to all given coordinates of less than 10000:", size)
}

func parseFile(name string) ([][2]int, error) {
	input, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(input), "\n")

	var coords [][2]int
	for _, line := range lines {
		if line == "" {
			continue
		}

		var coord [2]int
		if _, err := fmt.Sscanf(line, "%d, %d", &coord[0], &coord[1]); err != nil {
			return nil, err
		}

		coords = append(coords, coord)
	}

	return coords, nil
}

package main

import (
	"fmt"
	"strconv"
)

const input = "681901"

func main() {
	inputN, _ := strconv.Atoi(input)

	scores := []int{3, 7}
	elf1 := 0
	elf2 := 1

	for i := 0; ; i++ {
		sum := scores[elf1] + scores[elf2]

		for _, ch := range strconv.Itoa(sum) {
			digit, _ := strconv.Atoi(string(ch))
			scores = append(scores, digit)
		}

		if len(scores) == inputN+10 {
			fmt.Print("scores of the ten recipes immediately after the number of recipes in your puzzle input: ")
			for i := inputN; i < inputN+10; i++ {
				fmt.Print(scores[i%len(scores)])
			}
			fmt.Println()
		}

		elf1 = (elf1 + 1 + scores[elf1]) % len(scores)
		elf2 = (elf2 + 1 + scores[elf2]) % len(scores)

		if i+len(input) <= len(scores) && join(scores[i:i+len(input)]) == input {
			fmt.Println("number of recipes that appear on the scoreboard to the left of the score sequence in your puzzle input:", i)
			break
		}
	}
}

func join(a []int) string {
	out := make([]byte, len(a))

	for i, v := range a {
		out[i] = byte(v + '0')
	}

	return string(out)
}

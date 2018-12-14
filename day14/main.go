package main

import (
	"fmt"
	"strconv"
)

const input = "681901"

func main() {
	inputN, _ := strconv.Atoi(input)

	scores := []byte("37")
	elf1 := 0
	elf2 := 1

	for i := 0; ; i++ {
		score1, _ := strconv.Atoi(string(scores[elf1]))
		score2, _ := strconv.Atoi(string(scores[elf2]))
		sum := score1 + score2

		for _, ch := range strconv.Itoa(sum) {
			scores = append(scores, byte(ch))
		}

		if len(scores) == inputN+10 {
			fmt.Print("scores of the ten recipes immediately after the number of recipes in your puzzle input: ")
			for i := inputN; i < inputN+10; i++ {
				fmt.Print(string(scores[i%len(scores)]))
			}
			fmt.Println()
		}

		elf1 = (elf1 + 1 + score1) % len(scores)
		elf2 = (elf2 + 1 + score2) % len(scores)

		if i+len(input) <= len(scores) && string(scores[i:i+len(input)]) == input {
			fmt.Println("number of recipes that appear on the scoreboard to the left of the score sequence in your puzzle input:", i)
			break
		}
	}
}

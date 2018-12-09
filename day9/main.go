package main

import (
	"container/ring"
	"fmt"
	"os"
)

func play(players, lastValue int) int {
	scores := make([]int, players)
	circle := ring.New(1)
	circle.Value = 0

	for marble := 1; marble <= lastValue; marble++ {
		if marble%23 == 0 {
			circle = circle.Move(-7)
			scores[marble%players] += marble + circle.Value.(int)

			circle = circle.Prev()
			circle.Unlink(1)
			circle = circle.Next()

			continue
		}

		new := ring.New(1)
		new.Value = marble

		circle = new.Link(circle.Next().Next())
	}

	var highscore int
	for _, v := range scores {
		if v > highscore {
			highscore = v
		}
	}

	return highscore
}

func main() {
	var players, lastValue int

	f, _ := os.Open("input.txt")
	fmt.Fscanf(f, "%d players; last marble is worth %d points", &players, &lastValue)

	fmt.Println("winning Elf's score:", play(players, lastValue))
	fmt.Println("new winning Elf's score be if the number of the last marble were 100 times larger:", play(players, lastValue*100))

}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type cartMap [][]cell

type cell struct {
	ch                byte
	orig              byte
	origPos           [2]int
	intersectionCount int
}

func isCart(ch byte) bool {
	return ch == '>' || ch == '<' || ch == '^' || ch == 'v'
}

func flip(ch byte) byte {
	switch ch {
	case '>':
		return '<'
	case '<':
		return '>'
	case '^':
		return 'v'
	case 'v':
		return '^'
	}

	panic("flip on non-cart!")
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	track := bytes.Split(f, []byte("\n"))

	var cartCount int
	m := make(cartMap, len(track))
	for i, v := range track {
		m[i] = make([]cell, len(v))
		for j, b := range v {
			m[i][j] = cell{ch: b, origPos: [2]int{j, i}}

			orig := b
			if isCart(b) {
				cartCount++
				switch b {
				case '>', '<':
					orig = '-'
				case '^', 'v':
					orig = '|'
				}
			}

			m[i][j].orig = orig
		}
	}

	var collided bool
	for tick := 0; ; tick++ {
		moved := make(map[[2]int]bool)

		for y := range m {
			for x, cell := range m[y] {
				if moved[cell.origPos] {
					continue
				}

				if isCart(cell.ch) {
					var dirX, dirY int
					switch cell.ch {
					case '>':
						dirX++
					case '<':
						dirX--
					case '^':
						dirY--
					case 'v':
						dirY++
					}

					nextY, nextX := y+dirY, x+dirX
					next := m[nextY][nextX]

					if isCart(next.ch) {
						if !collided {
							fmt.Printf("collision at (x,y): (%d,%d) on tick %d\n", nextX, nextY, tick)
						}

						collided = true
						m[nextY][nextX].ch = next.orig
						m[y][x].ch = cell.orig
						cartCount -= 2
						continue
					}

					switch next.ch {
					case '/', '\\':
						switch cell.ch {
						case '>':
							cell.ch = '^'
						case '<':
							cell.ch = 'v'
						case '^':
							cell.ch = '>'
						case 'v':
							cell.ch = '<'
						}

						if next.ch == '\\' {
							cell.ch = flip(cell.ch)
						}
					case '+':
						intersectionIndex := cell.intersectionCount % 3
						if intersectionIndex == 0 || intersectionIndex == 2 {
							switch cell.ch {
							case '<':
								cell.ch = 'v'
							case '>':
								cell.ch = '^'
							case '^':
								cell.ch = '<'
							case 'v':
								cell.ch = '>'
							}

							if intersectionIndex == 2 {
								cell.ch = flip(cell.ch)
							}
						}

						cell.intersectionCount++
					}

					next.ch = cell.ch
					cell.ch = cell.orig

					next.origPos = cell.origPos
					cell.origPos = [2]int{}

					next.intersectionCount = cell.intersectionCount
					cell.intersectionCount = 0

					m[y][x] = cell
					m[nextY][nextX] = next
					moved[next.origPos] = true
				}

			}
		}

		if cartCount == 1 {
			for y, v := range m {
				for x, b := range v {
					if isCart(b.ch) {
						fmt.Printf("last surviving cart at (x,y): (%d,%d)\n", x, y)
					}
				}
			}

			break
		}
	}
}

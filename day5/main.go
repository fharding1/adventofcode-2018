package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"unicode"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// This is 10x faster than strings.ToLower or bytes.ToLower. Shortened the
// execution time from 6s to 600ms :exploding_head:.
func toLower(ch byte) byte {
	return byte(unicode.ToLower(rune(ch)))
}

func react(polymer []byte) []byte {
	changed := true

	for changed {
		oldChecksum := crc32.Checksum(polymer, crc32.IEEETable)

		for i := range polymer {
			cur := polymer[i]

			if i-1 >= 0 {
				prev := polymer[i-1]

				if prev != cur && toLower(prev) == toLower(cur) {
					polymer = append(polymer[:i-1], polymer[i+1:]...)
					break
				}
			}

			if i+1 < len(polymer) {
				next := polymer[i+1]

				if next != cur && toLower(next) == toLower(cur) {
					polymer = append(polymer[:i], polymer[i+2:]...)
					break
				}
			}
		}

		if oldChecksum == crc32.Checksum(polymer, crc32.IEEETable) {
			changed = false
		}
	}

	return polymer
}

func main() {
	polymer, _ := ioutil.ReadFile("input.txt")
	polymer = bytes.Trim(polymer, "\n")

	polymer = react(polymer)

	fmt.Println("polymer length after reacting:", len(polymer))

	shortestPolymerLen := len(polymer)
	chars := []byte(alphabet)

	for _, char := range chars {
		testPolymer := make([]byte, len(polymer))
		copy(testPolymer, polymer)

		testPolymer = bytes.Replace(testPolymer, []byte{char}, []byte{}, -1)
		testPolymer = bytes.Replace(testPolymer, bytes.ToUpper([]byte{char}), []byte{}, -1)

		testPolymer = react(testPolymer)

		if len(testPolymer) < shortestPolymerLen {
			shortestPolymerLen = len(testPolymer)
		}
	}

	fmt.Println("shortest polymer length:", shortestPolymerLen)
}

package main

import (
	"aoc2022"
	"fmt"
)

func main() {
	timer := aoc2022.Time()
	defer timer()

	input := aoc2022.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	sequenceLength := 14

	current := []rune(lines[0][:sequenceLength])
	for i, token := range lines[0][sequenceLength:] {
		index := i % sequenceLength

		if len(aoc2022.Set(current)) == sequenceLength {
			return i + sequenceLength
		}
		current[index] = token
	}

	return -1
}

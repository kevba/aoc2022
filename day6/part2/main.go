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
	sequence := 0
	sequenceLength := 14

	current := []rune(lines[0][:sequenceLength])
	for i, token := range lines[0][sequenceLength:] {
		if uniqueCounter(current) == sequenceLength {
			return i + sequenceLength
		}
		index := i % sequenceLength
		current[index] = token
	}

	return sequence
}

func uniqueCounter(slice []rune) int {
	counter := map[rune]int{}

	for _, token := range slice {
		if entry, ok := counter[token]; ok {
			entry++
		} else {
			counter[token] = 1
		}
	}

	return len(counter)
}

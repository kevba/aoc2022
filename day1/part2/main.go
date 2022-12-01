package main

import (
	"aoc2022"
	"fmt"
	"sort"
)

func main() {
	timer := aoc2022.Time()
	defer timer()

	input := aoc2022.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	elves := getElvesCalories(lines)
	sort.Ints(elves)

	last := len(elves) - 1
	answer := elves[last] + elves[last-1] + elves[last-2]

	return answer
}

func getElvesCalories(lines []string) []int {
	elves := []int{0}

	for _, line := range lines {
		if line == "" {
			elves = append(elves, 0)
		} else {
			elves[len(elves)-1] += aoc2022.Atoi(line)
		}
	}

	return elves
}

package main

import (
	"aoc2022"
	"fmt"
	"strings"
)

func main() {
	timer := aoc2022.Time()
	defer timer()

	input := aoc2022.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	counter := 0

	for _, l := range lines {
		sections := strings.Split(l, ",")
		elf1 := getElfSections(sections[0])
		elf2 := getElfSections(sections[1])

		if isRangeOverlap(elf1, elf2) {
			counter += 1
		}
	}

	return counter
}

func getElfSections(sectionRange string) [2]int {
	sections := strings.Split(sectionRange, "-")
	return [2]int{
		aoc2022.Atoi(sections[0]),
		aoc2022.Atoi(sections[1]),
	}
}

func isRangeOverlap(r1, r2 [2]int) bool {
	if r1[0] <= r2[0] && r1[1] >= r2[0] {
		return true
	} else if r2[0] <= r1[0] && r2[1] >= r1[0] {
		return true
	}

	return false
}

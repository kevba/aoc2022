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

		if elf1[0] <= elf2[0] && elf1[1] >= elf2[1] {
			counter += 1
		} else if elf2[0] <= elf1[0] && elf2[1] >= elf1[1] {
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

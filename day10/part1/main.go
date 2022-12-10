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
	solution := 0

	cycle := 1
	x := 1

	signalCheck := 20
	signalInterval := 40

	nextCycle := func() {
		if cycle == signalCheck {
			solution += x * cycle
			signalCheck += signalInterval
		}
		cycle++

	}

	for _, line := range lines {
		parsed := strings.Split(line, " ")

		nextCycle()
		switch parsed[0] {
		case "noop":
		case "addx":
			nextCycle()
			x += aoc2022.Atoi(parsed[1])

		}
	}

	return solution
}

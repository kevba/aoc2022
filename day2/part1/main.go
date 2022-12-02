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
	score := 0

	for _, l := range lines {
		moves := strings.Split(l, " ")

		you := mapHelper[moves[1]]
		opp := mapHelper[moves[0]]

		// Always add the pick score
		score += you

		// Draw
		if you == opp {
			score += 3
		}

		// Wins
		if you == mapHelper["X"] && opp == mapHelper["C"] {
			score += 6
		}
		if you == mapHelper["Y"] && opp == mapHelper["A"] {
			score += 6
		}
		if you == mapHelper["Z"] && opp == mapHelper["B"] {
			score += 6
		}
	}

	return score
}

var mapHelper = map[string]int{
	// ROCK
	"A": 1,
	"X": 1,
	// PAPER
	"B": 2,
	"Y": 2,
	// Scissors
	"C": 3,
	"Z": 3,
}

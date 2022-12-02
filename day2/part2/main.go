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

		you := 0
		opponent := mapHelper[moves[0]]

		// Lose
		if moves[1] == "X" {
			switch opponent {
			case 1:
				you = 3
			case 2:
				you = 1
			case 3:
				you = 2
			}
		}
		// Draw
		if moves[1] == "Y" {
			you = opponent
			score += 3
		}
		// Win
		if moves[1] == "Z" {
			switch opponent {
			case 1:
				you = 2
			case 2:
				you = 3
			case 3:
				you = 1
			}
			score += 6
		}

		// Always add the pick score
		score += you

	}

	return score
}

var mapHelper = map[string]int{
	// ROCK
	"A": 1,
	// PAPER
	"B": 2,
	// Scissors
	"C": 3,
}

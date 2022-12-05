package main

import (
	"aoc2022"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	timer := aoc2022.Time()
	defer timer()

	input := aoc2022.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) string {
	sequence := ""
	inputSplitIndex := getMoveIndex(lines)

	state := parseInitialState(lines[:inputSplitIndex])
	moves := parseMoves(lines[inputSplitIndex+1:])

	for _, m := range moves {
		topIndex := len(state[m.from]) - m.amount
		toMove := state[m.from][topIndex:]

		state[m.from] = state[m.from][:topIndex]
		state[m.to] = append(state[m.to], toMove...)
	}

	for _, stack := range state {
		sequence += string(stack[len(stack)-1])
	}

	return sequence
}

func getMoveIndex(lines []string) int {
	for i, l := range lines {
		if strings.Contains(l, " 1") {
			return i + 1
		}
	}
	return -1
}

func parseInitialState(lines []string) [][]byte {
	stacks := make([][]byte, 0)
	for _, l := range lines {
		for i := 1; i < len(l); i += 4 {
			index := (i - 1) / 4
			c := l[i]

			if len(stacks)-1 < index {
				stacks = append(stacks, []byte{})
			}

			if c != ' ' {
				stacks[index] = append([]byte{c}, stacks[index]...)
			}
		}
	}

	return stacks
}
func parseMoves(lines []string) []move {
	moves := []move{}
	r, _ := regexp.Compile(`([0-9]{1,})`)
	for _, l := range lines {
		results := r.FindAllString(l, -1)
		move := move{
			amount: aoc2022.Atoi(results[0]),
			from:   aoc2022.Atoi(results[1]) - 1,
			to:     aoc2022.Atoi(results[2]) - 1,
		}
		moves = append(moves, move)
	}

	return moves
}

type move struct {
	amount int
	from   int
	to     int
}

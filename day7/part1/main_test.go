package main

import (
	"aoc2022"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2022.GetTestInput()
	solution := 95437

	answer := solve(input)
	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}

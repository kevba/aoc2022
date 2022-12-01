package main

import (
	"aoc2022"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2022.GetTestInput()

	solution := 24000
	answer := solve(input)

	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}

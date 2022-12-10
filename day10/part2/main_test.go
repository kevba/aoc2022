package main

import (
	"aoc2022"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2022.GetTestInput()
	solution := "REHPRLUB"

	answer := solve(input)
	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}

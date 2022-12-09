package main

import (
	"aoc2022"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2022.ReadInputFile("../input_test_part_2.txt")
	solution := 36

	answer := solve(input)
	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}

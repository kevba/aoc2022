package main

import (
	"aoc2022"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2022.GetTestInput()
	solution := 157

	answer := solve(input)
	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}

func TestNewRucksack(t *testing.T) {
	input := aoc2022.GetTestInput()

	for _, l := range input {
		r := newRucksack(l)

		if len(r.comp1) != len(r.comp2) {
			t.Errorf("compartements are not equal length \n rucksack: %v \n 1: %v \n 2: %v", l, r.comp1, r.comp2)
		}
	}
}

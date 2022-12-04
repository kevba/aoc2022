package main

import (
	"aoc2022"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	timer := aoc2022.Time()
	defer timer()

	input := aoc2022.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	prioritySum := 0

	for _, l := range lines {
		r := newRucksack(l)
		prioritySum += r.getPriorityOfDouble()
	}

	return prioritySum
}

type rucksack struct {
	comp1 string
	comp2 string
}

func (r *rucksack) getPriorityOfDouble() int {
	for _, item := range r.comp1 {
		if strings.ContainsRune(r.comp2, item) {
			priority := strings.IndexRune(alphabet, unicode.ToUpper(item)) + 1
			if unicode.IsUpper(item) {
				priority += 26
			}

			return priority
		}
	}

	return 0
}

func newRucksack(contents string) *rucksack {
	middle := len(contents) / 2
	comp1 := contents[:middle]
	comp2 := contents[middle:]

	return &rucksack{
		comp1: comp1,
		comp2: comp2,
	}
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

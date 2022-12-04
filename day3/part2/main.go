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

	for i := 0; i < len(lines); i += 3 {
		elf1 := lines[i]
		elf2 := lines[i+1]
		elf3 := lines[i+2]

		for _, item := range elf1 {
			if strings.ContainsRune(elf2, item) {
				if strings.ContainsRune(elf3, item) {
					prioritySum += getPriority(item)
					break
				}
			}
		}

	}

	return prioritySum
}

func getPriority(r rune) int {
	priority := strings.IndexRune(alphabet, unicode.ToUpper(r)) + 1
	if unicode.IsUpper(r) {
		priority += 26
	}
	return priority

}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func solve(lines []string) string {
	hSync := 40
	vSync := 240 / 40

	cycle := 0
	spritePos := 1

	screen := setupScreen(vSync, hSync)

	endCycle := func() {
		if aoc2022.IntAbs((cycle%hSync)-spritePos) <= 1 {
			screen[cycle/hSync][cycle%hSync] = "#"
		}

		cycle++
	}

	for _, line := range lines {
		endCycle()
		parsed := strings.Split(line, " ")

		switch parsed[0] {
		case "noop":
		case "addx":
			endCycle()
			spritePos += aoc2022.Atoi(parsed[1])

		}
	}

	printScreen(screen)
	return "REHPRLUB"
}

func setupScreen(v, h int) [][]string {
	screen := [][]string{}
	for i := 0; i < v; i++ {
		row := make([]string, h)
		for j := range row {
			row[j] = "."
		}
		screen = append(screen, row)
	}
	return screen
}

func printScreen(screen [][]string) {
	for _, r := range screen {
		for _, v := range r {
			fmt.Print(v)
		}
		fmt.Print("\n")
	}
}

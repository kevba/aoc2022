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
	trees := make([][]int, len(lines))
	for i, line := range lines {
		trees[i] = aoc2022.AtoiSlice(strings.Split(line, ""))
	}

	visible := make([][]int, len(lines))
	for i := 0; i < len(visible); i++ {
		visible[i] = make([]int, len(lines[0]))
	}

	for y, treeRow := range trees {
		for x, _ := range treeRow {
			currentTree := trees[y][x]
			if visible[y][x] == 1 {
				continue
			}

			right := isHighest(trees[y][x+1:], currentTree)
			left := isHighest(trees[y][:x], currentTree)

			col := []int{}
			for j := range trees {
				col = append(col, trees[j][x])
			}

			top := isHighest(col[:y], currentTree)
			bottom := isHighest(col[y+1:], currentTree)

			if left || right || top || bottom {
				visible[y][x] = 1
			}

		}
	}

	counter := 0
	for _, r := range visible {
		for _, v := range r {
			counter += v
		}
	}

	return counter
}

func isHighest(entries []int, value int) bool {
	for _, v := range entries {
		if value <= v {
			return false
		}
	}

	return true
}

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

	currentBestScore := 0

	for y, treeRow := range trees {
		for x, _ := range treeRow {
			currentTree := trees[y][x]

			right := getSeeingDistance(trees[y][x+1:], currentTree)
			left := getSeeingDistance(aoc2022.ReverseSlice(trees[y][:x]), currentTree)

			col := []int{}
			for j := range trees {
				col = append(col, trees[j][x])
			}

			top := getSeeingDistance(aoc2022.ReverseSlice(col[:y]), currentTree)
			bottom := getSeeingDistance(col[y+1:], currentTree)

			scenicScore := left * right * top * bottom

			if scenicScore > currentBestScore {
				currentBestScore = scenicScore
			}

		}
	}

	return currentBestScore
}

func getSeeingDistance(entries []int, value int) int {
	for i, v := range entries {
		if value <= v {
			return i + 1
		}
	}

	return len(entries)
}

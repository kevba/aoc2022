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
	head := []int{0, 0}
	tails := [][]int{}
	for i := 0; i < 9; i++ {
		tails = append(tails, []int{0, 0})
	}

	visited := map[string]int{}

	for _, move := range lines {
		parts := strings.Split(move, " ")
		direction := parts[0]
		amount := aoc2022.Atoi(parts[1])

		for i := 0; i < amount; i++ {
			switch direction {
			case "L":
				head[0] -= 1
			case "R":
				head[0] += 1
			case "U":
				head[1] += 1
			case "D":
				head[1] -= 1
			}

			tails[0] = updateTail(head, tails[0])
			for i, t := range tails[:len(tails)-1] {
				tails[i+1] = updateTail(t, tails[i+1])
			}

			lastTail := tails[len(tails)-1]
			visited[fmt.Sprintf("%v:%v", lastTail[0], lastTail[1])] = 1
		}
	}

	return len(visited)
}

func updateTail(head []int, tail []int) []int {
	xDelta := tail[0] - head[0]
	yDelta := tail[1] - head[1]

	newTail := []int{
		tail[0], tail[1],
	}

	// No longer touching
	if aoc2022.IntAbs(xDelta) > 1 {
		newTail[0] = head[0] + xDelta/aoc2022.IntAbs(xDelta)

		if aoc2022.IntAbs(yDelta) == 1 {
			newTail[1] = head[1]
		}
	}

	if aoc2022.IntAbs(yDelta) > 1 {
		newTail[1] = head[1] + yDelta/aoc2022.IntAbs(yDelta)

		if aoc2022.IntAbs(xDelta) == 1 {
			newTail[0] = head[0]
		}
	}

	return newTail
}

func printVisited(visited map[string]int) {
	const size = 25

	for y := size; y > -size; y-- {
		for x := -size; x < size; x++ {
			if _, ok := visited[fmt.Sprintf("%v:%v", x, y)]; ok {
				if fmt.Sprintf("%v:%v", x, y) == "0:0" {

					fmt.Print("s")
				} else {
					fmt.Print("#")
				}
			} else {
				fmt.Print(".")

			}
		}
		fmt.Println("")
	}
}

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
	tail := []int{0, 0}

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

			xDelta := aoc2022.IntAbs(tail[0] - head[0])
			yDelta := aoc2022.IntAbs(tail[1] - head[1])

			// No longer touching
			if xDelta > 1 || yDelta > 1 {
				switch direction {
				case "R":
					if yDelta == 0 {
						tail[0] = head[0] - 1
					} else if yDelta > 0 {
						tail[0] = head[0] - 1
						tail[1] = head[1]
					}
				case "L":
					if yDelta == 0 {
						tail[0] = head[0] + 1
					} else if yDelta > 0 {
						tail[0] = head[0] + 1
						tail[1] = head[1]
					}
				case "U":
					if xDelta == 0 {
						tail[1] = head[1] - 1
					} else if yDelta > 0 {
						tail[1] = head[1] - 1
						tail[0] = head[0]
					}
				case "D":
					if xDelta == 0 {
						tail[1] = head[1] + 1
					} else if yDelta > 0 {
						tail[1] = head[1] + 1
						tail[0] = head[0]
					}
				}

			}

			visited[fmt.Sprintf("%v:%v", tail[0], tail[1])] = 1
		}
	}

	return len(visited)
}

func printVisited(visited map[string]int) {
	const size = 10

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

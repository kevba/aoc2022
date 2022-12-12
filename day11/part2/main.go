package main

import (
	"aoc2022"
	"fmt"
	"regexp"
)

func main() {
	timer := aoc2022.Time()
	defer timer()

	input := aoc2022.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	monkeys := parse(lines)

	divider := 1
	for _, m := range monkeys {
		divider = divider * m.testVal
	}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			items := m.inspectAndThrow(divider)
			for k, v := range items {
				monkeys[k].catch(v)
			}
		}
	}

	highest := 0
	second := 0
	for _, m := range monkeys {
		if highest <= m.inspectCounter {
			second = highest
			highest = m.inspectCounter
		} else if second <= m.inspectCounter {
			second = m.inspectCounter
		}
	}

	return highest * second
}

type monkey struct {
	items     []int
	operation func(v int) int
	testVal   int
	testTrue  int
	testFalse int

	inspectCounter int
}

func (m *monkey) catch(items []int) {
	m.items = append(m.items, items...)
}

func (m *monkey) inspectAndThrow(divider int) map[int][]int {

	items := aoc2022.ReverseSlice(m.items)
	m.items = []int{}

	actions := map[int][]int{
		m.testTrue:  {},
		m.testFalse: {},
	}

	for _, item := range items {
		m.inspectCounter++
		item = m.operation(item) % divider

		if item%m.testVal == 0 {
			actions[m.testTrue] = append(actions[m.testTrue], item)
		} else {
			actions[m.testFalse] = append(actions[m.testFalse], item)
		}
	}

	return actions
}

func parse(lines []string) []*monkey {
	monkeys := []*monkey{}
	r, _ := regexp.Compile(`([0-9]{1,})`)
	operationR, _ := regexp.Compile(`([-+/*]{1})`)

	for i := 0; i < len(lines); i += 7 {
		items := aoc2022.AtoiSlice(r.FindAllString(lines[i+1], -1))

		operationOperator := operationR.FindString(lines[i+2])
		operationVal := r.FindString(lines[i+2])

		testVal := aoc2022.Atoi(r.FindString(lines[i+3]))
		testTrue := aoc2022.Atoi(r.FindString(lines[i+4]))
		testFalse := aoc2022.Atoi(r.FindString(lines[i+5]))

		operation := func(v int) int { return v }

		switch operationOperator {
		case "*":
			if operationVal != "" {
				operation = func(v int) int { return v * aoc2022.Atoi(operationVal) }
			} else {
				operation = func(v int) int { return v * v }
			}
		case "+":
			if operationVal != "" {
				operation = func(v int) int { return v + aoc2022.Atoi(operationVal) }
			} else {
				operation = func(v int) int { return v + v }
			}
		}

		m := &monkey{
			items:          items,
			operation:      operation,
			testVal:        testVal,
			testTrue:       testTrue,
			testFalse:      testFalse,
			inspectCounter: 0,
		}

		monkeys = append(monkeys, m)

	}

	return monkeys
}

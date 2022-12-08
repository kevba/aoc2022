package aoc2022

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

var inputFile = "../input.txt"
var testInputFile = "../input_test.txt"

// GetInput returns the puzzle input in rows.
func GetInput() []string {
	return ReadInputFile(inputFile)
}

// GetInput returns the puzzle input in rows.
func GetTestInput() []string {
	return ReadInputFile(testInputFile)
}

func ReadInputFile(path string) []string {
	values := []string{}
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("cannot open %v : %v", path, err)
	}

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		values = append(values, line)
	}

	return values
}

// Atoi converts a string to an int. It fatals when this is not possible.
func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot convert %v to int: %v", s, err)
	}
	return num
}

// AtoiSlice converts a string slice to an intSlice. It fatals when this is not possible.
func AtoiSlice(s []string) []int {
	ints := []int{}

	for _, v := range s {
		ints = append(ints, Atoi(v))
	}

	return ints
}

func IntAbs(n int) int {
	return int(math.Abs(float64(n)))
}

func Time() func() {
	start := time.Now()
	return func() {
		log.Printf("solved in %v \n", time.Since(start))
	}
}

func Set[T comparable](s []T) []T {
	new := []T{}

	for _, v := range s {
		if !Contains(new, v) {
			new = append(new, v)
		}
	}

	return new
}

func Contains[T comparable](s []T, e T) bool {
	return FindIndex(s, e) != -1
}

func FindIndex[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

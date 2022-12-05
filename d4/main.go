package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task: https://adventofcode.com/2022/day/4

type Pair []int

func main() {
	input := readInput("input.txt")
	fullOverlaps := 0
	partialOverlaps := 0

	for _, ln := range input {
		full, partial := getOverlaps(ln)

		if full {
			fullOverlaps++
		}

		if partial {
			partialOverlaps++
		}
	}

	fmt.Println("Overlaps", fullOverlaps, partialOverlaps)
}

func getOverlaps(s string) (bool, bool) {
	var pairs []Pair

	p := strings.Split(s, ",")

	for _, v := range p {

		r := strings.Split(v, "-")
		var newRange Pair

		for _, v2 := range r {

			q, _ := strconv.Atoi(v2)
			newRange = append(newRange, q)
		}

		pairs = append(pairs, newRange)
	}

	full := getFull(pairs[0], pairs[1])
	partial := getPartial(pairs[0], pairs[1])

	return full, partial
}

func getFull(a, b Pair) bool {
	if a[0] <= b[0] && a[1] >= b[1] {
		return true
	}

	if b[0] <= a[0] && b[1] >= a[1] {
		return true
	}

	return false
}

func getPartial(a, b Pair) bool {
	if a[1] >= b[0] && a[0] <= b[1] {
		return true
	}

	return false
}

func readInput(f string) []string {
	bs, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\n")
}

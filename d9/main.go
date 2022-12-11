package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/9

var h = [2]int{0, 0}
var t = [2]int{0, 0}
var positions = map[string]int{"0,0": 1}

func main() {
	input := readInput("input.txt")

	for _, ln := range input {
		p := strings.Split(ln, " ")

		move(p[0], toInt(p[1]))
	}

	// Part 1
	fmt.Println("Positions tail visited", len(positions))
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return i
}

func moveHead(dir string) {
	switch dir {
	case "U":
		h[1]++
	case "D":
		h[1]--
	case "R":
		h[0]++
	case "L":
		h[0]--
	}
}

func moveTail(dir string) {
	switch dir {
	case "U":
		if h[1]-t[1] > 1 {

			t[1]++

			if t[0] != h[0] {
				t[0] = h[0]
			}

			positions[fmt.Sprintf("%v,%v", t[0], t[1])] += 1
		}

	case "D":
		if t[1]-h[1] > 1 {

			t[1]--

			if t[0] != h[0] {
				t[0] = h[0]
			}

			positions[fmt.Sprintf("%v,%v", t[0], t[1])] += 1
		}

	case "R":
		if h[0]-t[0] > 1 {

			t[0]++

			if t[1] != h[1] {
				t[1] = h[1]
			}

			positions[fmt.Sprintf("%v,%v", t[0], t[1])] += 1
		}

	case "L":
		if t[0]-h[0] > 1 {

			t[0]--

			if t[1] != h[1] {
				t[1] = h[1]
			}

			positions[fmt.Sprintf("%v,%v", t[0], t[1])] += 1
		}
	}

}

func move(dir string, dist int) {
	for i := 0; i < dist; i++ {
		moveHead(dir)
		moveTail(dir)
	}
}

func readInput(f string) []string {
	bs, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\n")
}

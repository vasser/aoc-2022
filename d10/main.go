package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/10

var cycle = 0
var maxCycle = 239
var slice = 40

var acc = 1
var o = []int{}

var m = [6][40]string{
	{"#", "#", "#", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
}
var curPos = 0
var maxPos = 39

func main() {
	input := readInput("input.txt")

	for _, ln := range input {
		if ln == "noop" {
			handleCycle(&cycle)
		} else {
			i := strings.Split(ln, " ")
			val := toInt(i[1])

			handleCycle(&cycle)

			handleCycle(&cycle)

			acc += val
		}
	}

	// Task 1
	sum := 0
	for _, v := range o {
		sum += v
	}
	fmt.Println("Task 1", sum)

	// Task 2
	fmt.Println("Task 2")
	for _, v := range m {
		for _, vv := range v {
			fmt.Print(vv)
		}
		fmt.Print("\n")
	}
}

func handleCycle(cycle *int) {
	*cycle++
	if (*cycle-20)%slice == 0 {
		o = append(o, acc**cycle)
	}

	row := math.Floor(float64(*cycle / slice))

	if *cycle > maxCycle {
		return
	}

	spriteStart := acc - 1
	spriteEnd := acc + 1
	val := "."

	if (curPos >= spriteStart) && (curPos <= spriteEnd) {
		val = "#"
	}

	m[int(row)][curPos] = val

	if curPos == maxPos {
		curPos = 0
	} else {
		curPos++
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return i
}

func readInput(f string) []string {
	bs, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\n")
}

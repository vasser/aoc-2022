package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Task: https://adventofcode.com/2022/day/5

type Stack []string

var c1 = Stack{"H", "C", "R"}
var c2 = Stack{"B", "J", "H", "L", "S", "F"}
var c3 = Stack{"R", "M", "D", "H", "J", "T", "Q"}
var c4 = Stack{"S", "G", "R", "H", "Z", "B", "J"}
var c5 = Stack{"R", "P", "F", "Z", "T", "D", "C", "B"}
var c6 = Stack{"T", "H", "C", "G"}
var c7 = Stack{"S", "N", "V", "Z", "B", "P", "W", "L"}
var c8 = Stack{"R", "J", "Q", "G", "C"}
var c9 = Stack{"L", "D", "T", "R", "H", "P", "F", "S"}

var crates = []Stack{c1, c2, c3, c4, c5, c6, c7, c8, c9}

func main() {
	input := readInput("input.txt")

	for _, line := range input {
		re := regexp.MustCompile(`[0-9]+`)
		res := re.FindAllString(line, -1)

		num, _ := strconv.Atoi(res[0])
		f, _ := strconv.Atoi(res[1])
		f = f - 1
		t, _ := strconv.Atoi(res[2])
		t = t - 1

		cutIdx := 0
		if len(crates[f])-num > 0 {
			cutIdx = len(crates[f]) - num
		}

		temp := crates[f][cutIdx:]

		// task 1
		// for i := len(temp) - 1; i >= 0; i-- {
		// 	crates[t] = append(crates[t], temp[i])
		// 	crates[f] = crates[f][:len(crates[f])-1]
		// }

		// task 2
		crates[t] = append(crates[t], temp...)
		crates[f] = crates[f][:len(crates[f])-len(temp)]

	}

	for i, _ := range crates {
		fmt.Print(crates[i][len(crates[i])-1])
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

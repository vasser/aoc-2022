package main

import (
	"fmt"
	"os"
	"strings"
)

// Task: https://adventofcode.com/2022/day/6

var p_shift = 4
var m_shift = 14

func isUnique(s string) bool {
	m := make(map[string]int)

	for _, c := range s {
		m[string(c)]++

		if m[string(c)] > 1 {
			return false
		}
	}

	return true
}

func main() {
	input := readInput("input.txt")

	position := 0

	var p string
	var m string

	found_p := 0
	found_m := 0

	for {
		if found_p == 0 {
			p = strings.Join(input[position:position+p_shift], "")

			if isUnique(p) {
				found_p = position + p_shift
			}
		}

		if found_m == 0 {
			m = strings.Join(input[position:position+m_shift], "")

			if isUnique(m) {
				found_m = position + m_shift
			}
		}

		if found_p > 0 && found_m > 0 {
			break
		}

		position++
	}

	fmt.Println(found_p, found_m)
}

func readInput(f string) []string {
	bs, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "")
}

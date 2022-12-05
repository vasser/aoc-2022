package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Task: https://adventofcode.com/2022/day/1

func main() {
	input := readInput("input.txt")

	max := make([]int, 3)
	total := 0
	tmp := 0

	for _, val := range input {
		if val == "" {
			if tmp > max[0] {
				max[0] = tmp
				sort.Ints(max)
			}
			tmp = 0
			continue
		}

		i, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("could not convert", val)
		}

		tmp = tmp + i

	}

	for _, val := range max {
		total = total + val
	}

	fmt.Println("Max vals are", max)
	fmt.Println("Top 3 are", total)
}

func readInput(f string) []string {
	bs, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\n")
}

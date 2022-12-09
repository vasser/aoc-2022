package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/8

type Forest [][]int

var input = readInput("input.txt")

func main() {
	total := len(input)*2 + len(input[0])*2 - 4
	scenicScore := 0

	for i, s := range input {
		if i == 0 || i == len(input)-1 {
			continue
		}

		for j := range s {
			if j == 0 || j == len(s)-1 {
				continue
			}

			isVisible, score := calculateTree(i, j)

			if isVisible {
				total++
			}

			if score > scenicScore {
				scenicScore = score
			}
		}
	}

	fmt.Printf("Total visible: %v\n", total)
	fmt.Printf("Scenic score: %v\n", scenicScore)
}

func calculateTree(col, ln int) (bool, int) {
	val := input[col][ln]

	top, ts := isVisibleTop(val, col, ln)
	bottom, bs := isVisibleBottom(val, col, ln)
	left, ls := isVisibleLeft(val, col, ln)
	right, rs := isVisibleRight(val, col, ln)

	isVisible := top || bottom || left || right
	score := ts * bs * ls * rs

	return isVisible, score
}

func isVisibleTop(val, col, ln int) (bool, int) {
	isVisible := true
	score := 0

	for i := col - 1; i >= 0; i-- {

		score++

		if input[i][ln] >= val {
			isVisible = false
			break
		}
	}

	return isVisible, score
}

func isVisibleBottom(val, col, ln int) (bool, int) {
	isVisible := true
	score := 0

	for i := col + 1; i <= len(input)-1; i++ {

		score++

		if input[i][ln] >= val {
			isVisible = false
			break
		}
	}

	return isVisible, score
}

func isVisibleLeft(val, col, ln int) (bool, int) {
	isVisible := true
	score := 0

	for i := ln - 1; i >= 0; i-- {

		score++

		if input[col][i] >= val {
			isVisible = false
			break
		}
	}

	return isVisible, score
}

func isVisibleRight(val, col, ln int) (bool, int) {
	isVisible := true
	score := 0

	for i := ln + 1; i <= len(input[ln])-1; i++ {

		score++

		if input[col][i] >= val {
			isVisible = false
			break
		}
	}

	return isVisible, score
}

func readInput(f string) Forest {
	bs, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "\n")
	var i Forest

	for _, v := range s {
		var t []int

		for _, v := range strings.Split(v, "") {
			val, _ := strconv.Atoi(v)
			t = append(t, val)
		}

		i = append(i, t)
	}

	return i
}

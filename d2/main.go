package main

import (
	"fmt"
	"os"
	"strings"
)

// Task: https://adventofcode.com/2022/day/2

// Values
var m = map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

func main() {
	input := readInput("input.txt")
	score1 := 0
	score2 := 0

	for _, val := range input {
		score1 = score1 + getScore1(val)
		score2 = score2 + getScore2(val)
	}

	fmt.Println("Score:", score1, score2)
}

func getScore1(round string) int {
	// Win combinations
	r := map[string]int{
		"A Y": 6, "B Z": 6, "C X": 6, // wins
		"A X": 3, "B Y": 3, "C Z": 3, // ties
		"A Z": 0, "B X": 0, "C A": 0, // loses
	}

	pair := strings.Split(round, " ")

	return m[pair[1]] + r[round]
}

func getScore2(round string) int {
	// Win combinations
	r := map[string]int{
		"A Z": m["Y"] + 6, "B Z": m["Z"] + 6, "C Z": m["X"] + 6, // wins
		"A Y": m["X"] + 3, "B Y": m["Y"] + 3, "C Y": m["Z"] + 3, // ties
		"A X": m["Z"], "B X": m["X"], "C X": m["Y"], // loses
	}

	return r[round]
}

func readInput(f string) []string {
	bs, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\n")
}

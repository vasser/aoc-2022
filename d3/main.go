package main

import (
	"fmt"
	"os"
	"strings"
)

// Task: https://adventofcode.com/2022/day/3

var values = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func main() {
	input := readInput("input.txt")
	sum1 := 0
	sum2 := 0

	for i, ln := range input {
		// 1st task
		sum1 += getCommonItemInRucsackValue(ln)

		// 2nd task
		if i == 0 || i%3 == 0 {
			sum2 += getLabelValue(toSlice(input[i]), toSlice(input[i+1]), toSlice(input[i+2]))
		}
	}

	fmt.Println("Priorities", sum1, sum2)
}

func toSlice(s string) []string {
	return strings.Split(s, "")
}

func toMap(s []string) map[string]int {
	m := make(map[string]int)

	for _, v := range s {
		m[v] = 1
	}

	return m
}

func getValue(s string) int {
	if s == "" {
		return 0
	}

	return values[s]
}

func getCommonItemInRucsackValue(s string) int {
	sl := toSlice(s)
	half := len(sl) / 2
	c1 := sl[:half]
	c2 := sl[half:]

	d := getDuplicateInSlices(c1, c2)

	return getValue(d)
}

func getDuplicateInSlices(s1 []string, s2 []string) string {
	d := ""

	m := toMap(s1)

	for _, v := range s2 {
		if m[v] == 1 {
			d = v
			break
		}
	}

	return d
}

func getLabelValue(s1 []string, s2 []string, s3 []string) int {
	d := ""

	m := toMap(s1)

	for _, v := range s2 {
		if m[v] == 1 {
			for _, v2 := range s3 {
				if v == v2 {
					d = v
					break
				}
			}
		}
	}

	return getValue(d)
}

func readInput(f string) []string {
	bs, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\n")
}

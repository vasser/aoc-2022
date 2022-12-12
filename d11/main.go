package main

import (
	"fmt"
	"math"
	"sort"
)

// https://adventofcode.com/2022/day/11

type Op struct {
	o string
	v int
}

type Test struct {
	div int
	t   int
	f   int
}

type Monkey struct {
	items []int
	op    Op
	test  Test
}

func main() {
	round := 0

	result := []int{
		0, 0, 0, 0,
	}

	input := []Monkey{
		{
			items: []int{79, 98},
			op:    Op{"*", 19},
			test:  Test{23, 2, 3},
		},
		{
			items: []int{54, 65, 75, 74},
			op:    Op{"+", 6},
			test:  Test{19, 2, 0},
		},
		{
			items: []int{79, 60, 97},
			op:    Op{"*", -1},
			test:  Test{13, 1, 3},
		},
		{
			items: []int{74},
			op:    Op{"+", 3},
			test:  Test{17, 0, 1},
		},
	}

	l := len(input)

	for {
		i := round % l
		m := input[i]

		for _, v := range m.items {
			v = op(m.op.o, v, m.op.v)
			v = int(math.Floor(float64(v / 3)))

			if v%m.test.div == 0 {
				input[m.test.t].items = append(input[m.test.t].items, v)
			} else {
				input[m.test.f].items = append(input[m.test.f].items, v)
			}

			result[i]++
		}

		input[i].items = []int{}

		round++

		if round == l*20 {
			sort.Ints(result)
			fmt.Println("Task 1", result[len(result)-2]*result[len(result)-1])
			break
		}
	}
}

func op(op string, a, b int) int {
	if b == -1 {
		b = a
	}
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	default:
		return 0
	}
}

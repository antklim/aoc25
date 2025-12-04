package main

import (
	"fmt"
	"math"
	"os"
	"unicode"

	"github.com/antklim/aoc25/internal/utils"
)

func main() {
	joltages, err := utils.ReadInput("input.txt", mapJoltage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v", err)
		os.Exit(1)
	}

	maxJoltages := make([]int, 0, len(joltages))
	for _, j := range joltages {
		// maxJoltages = append(maxJoltages, maxJoltage(j))
		maxJoltages = append(maxJoltages, maxJoltage12(j))
	}

	sum := 0
	for _, j := range maxJoltages {
		sum += j
	}

	fmt.Printf("total joltages: %d\n", sum)
	os.Exit(0)
}

func mapJoltage(s string) ([]int, error) {
	result := make([]int, 0, len(s))
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return nil, fmt.Errorf("invalid input: %c is not a digit", c)
		}
		result = append(result, int(c-'0'))
	}
	return result, nil
}

func maxJoltage(a []int) int {
	var jolts []int

	for i, d := range a[:len(a)-1] {
		for _, v := range a[i+1:] {
			jolts = append(jolts, d*10+v)
		}
	}

	return max(jolts)
}

func maxJoltage12(a []int) int {
	result := 0
	pos := 0
	for i := range 12 {
		r := len(a) - 11 + i
		v, p := pickMaxAndPos(a[pos:r])
		result += int(math.Pow10(11-i)) * v
		pos += p + 1
	}
	return result
}

func pickMaxAndPos(a []int) (v int, pos int) {
	for i := range 10 {
		v = 9 - i
		if pos = firstXPos(a, v); pos != -1 {
			return
		}
	}

	panic("should pick a position earlier")
}

func max(a []int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func firstXPos(a []int, x int) int {
	for i, v := range a {
		if v == x {
			return i
		}
	}
	return -1
}

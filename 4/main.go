package main

import (
	"fmt"
	"os"

	"github.com/antklim/aoc25/internal/utils"
)

func main() {
	grid, err := utils.ReadInputFile("input.txt", mapInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v", err)
		os.Exit(1)
	}

	v := accessibleToForklift(grid)
	fmt.Printf("amount of rolls accessible by a forklift: %d\n", len(v))
	os.Exit(0)
}

func mapInput(s string) (string, error) {
	return s, nil
}

func isRoll(c rune) bool {
	return c == '@'
}

func accessibleToForklift(a []string) [][2]int {
	var result [][2]int

	for i, s := range a {
		for j, c := range s {
			if !isRoll(c) {
				continue
			}

			if numberOfadjacentRolls(a, i, j) < 4 {
				result = append(result, [2]int{i, j})
			}
		}
	}

	return result
}

func numberOfadjacentRolls(a []string, i, j int) int {
	result := 0
	// check row above
	if i > 0 {
		if j > 0 && isRoll(rune(a[i-1][j-1])) {
			result++
		}
		if isRoll(rune(a[i-1][j])) {
			result++
		}
		if j < len(a)-1 && isRoll(rune(a[i-1][j+1])) {
			result++
		}
	}

	// check left and right
	if j > 0 && isRoll(rune(a[i][j-1])) {
		result++
	}
	if j < len(a)-1 && isRoll(rune(a[i][j+1])) {
		result++
	}

	// check row bellow
	if i < len(a)-1 {
		if j > 0 && isRoll(rune(a[i+1][j-1])) {
			result++
		}
		if isRoll(rune(a[i+1][j])) {
			result++
		}
		if j < len(a)-1 && isRoll(rune(a[i+1][j+1])) {
			result++
		}
	}

	return result
}

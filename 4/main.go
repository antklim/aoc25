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

	// v := accessibleToForklift(grid)
	// fmt.Printf("amount of rolls accessible by a forklift: %d\n", len(v))
	v := removeAll(grid)
	fmt.Printf("amount of all remove rolls: %d\n", v)
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

func removeRolls(a []string, r [][2]int) []string {
	result := make([]string, 0, len(a))
	for _, s := range a {
		result = append(result, s)
	}

	for _, xy := range r {
		x, y := xy[0], xy[1]
		s := result[x]
		result[x] = s[:y] + "x" + s[y+1:]
	}

	return result
}

func removeAll(a []string) int {
	result := 0

	grid := a

	for {
		removableRolls := accessibleToForklift(grid)
		if len(removableRolls) == 0 {
			break
		}
		result += len(removableRolls)
		grid = removeRolls(grid, removableRolls)
	}

	return result
}

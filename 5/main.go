package main

import (
	"fmt"
	"os"
)

const (
	readingRanges = iota
	readingIngredientIDs
)

func main() {
	ranges, ids, err := readInputFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v", err)
		os.Exit(1)
	}

	v := filterFreshIDs(ids, ranges)
	fmt.Printf("amount of fresh IDs: %d\n", len(v))
	os.Exit(0)
}

func filterFreshIDs(ids []int, ranges [][2]int) []int {
	var result []int
	for _, id := range ids {
		if inRange(id, ranges) {
			result = append(result, id)
		}
	}
	return result
}

func inRange(id int, ranges [][2]int) bool {
	for _, v := range ranges {
		if id >= v[0] && id <= v[1] {
			return true
		}
	}
	return false
}

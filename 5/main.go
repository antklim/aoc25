package main

import (
	"fmt"
	"os"
	"slices"
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
	count := totalFreshIDs(ranges)
	fmt.Printf("total amount of fresh IDs: %d\n", count)
	os.Exit(0)
}

func filterFreshIDs(ids []int, ranges []Range) []int {
	var result []int
	for _, id := range ids {
		if inRange(id, ranges) {
			result = append(result, id)
		}
	}
	return result
}

func inRange(id int, ranges []Range) bool {
	for _, v := range ranges {
		if id >= v[0] && id <= v[1] {
			return true
		}
	}
	return false
}

type Range [2]int

func (r Range) Intersect(a Range) bool {
	return (r[0] >= a[0] && a[0] <= r[1]) || (r[0] >= a[1] && a[1] <= r[1])
}

func (r Range) IDs() []int {
	result := make([]int, 0, r[1]-r[0]+1)
	for i := r[0]; i <= r[1]; i++ {
		result = append(result, i)
	}
	return result
}

func (r Range) Contains(i int) bool {
	return r[0] <= i && i <= r[1]
}

func totalFreshIDs(ranges []Range) int {
	slices.SortFunc(ranges, func(a, b Range) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] == b[0] {
			return 0
		}
		return 1
	})

	from, to := -1, -1
	total := 0
	for _, r := range ranges {
		otherFrom, otherTo := r[0], r[1]
		if from < 0 || to < 0 {
			from = otherFrom
			to = otherTo
			total = to - from + 1
			continue
		}

		// next interval is within the previous
		if otherTo <= to {
			continue
		}

		if otherFrom <= to {
			from = to
			to = otherTo
			total += to - from
			continue
		}

		if otherFrom > to {
			from = otherFrom
			to = otherTo
			total += to - from + 1
			continue
		}
		panic("unreachable")
	}

	return total
}

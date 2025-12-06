package main

import (
	"strings"
	"testing"

	"github.com/antklim/aoc25/internal/utils"
)

const testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestReadInput(t *testing.T) {
	expectedRanges := [][2]int{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}
	expectedIDs := []int{1, 5, 8, 11, 17, 32}

	gotRanges, gotIDs, err := readInput(strings.NewReader(testInput))
	if err != nil {
		t.Errorf("failed to read input: %v", err)
	}
	if !utils.EqualSlices(expectedRanges, gotRanges) {
		t.Errorf("read input\nwant %v,\n got %v", expectedRanges, gotRanges)
	}
	if !utils.EqualSlices(expectedIDs, gotIDs) {
		t.Errorf("read input\nwant %v,\n got %v", expectedIDs, gotIDs)
	}
}

func TestFilterFreshIDs(t *testing.T) {
	ranges := [][2]int{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}
	ids := []int{1, 5, 8, 11, 17, 32}
	expected := []int{5, 11, 17}

	got := filterFreshIDs(ids, ranges)
	if !utils.EqualSlices(expected, got) {
		t.Errorf("fresh IDs are invalid\nwant %v,\n got %v", expected, got)
	}
}

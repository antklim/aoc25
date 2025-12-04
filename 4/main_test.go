package main

import (
	"strings"
	"testing"

	"github.com/antklim/aoc25/internal/utils"
)

const testInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestMapInput(t *testing.T) {
	expected := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}
	got, err := utils.ReadInput(strings.NewReader(testInput), mapInput)
	if err != nil {
		t.Errorf("failed to read input: %v", err)
	}
	if !utils.EqualSlices(expected, got) {
		t.Errorf("read input\nwant %v,\n got %v", expected, got)
	}
}

func TestAccessibleToForklift(t *testing.T) {
	expected := [][2]int{
		{0, 2}, {0, 3}, {0, 5}, {0, 6}, {0, 8},
		{1, 0},
		{2, 6},
		{4, 0}, {4, 9},
		{7, 0},
		{9, 0}, {9, 2}, {9, 8},
	}
	input, err := utils.ReadInput(strings.NewReader(testInput), mapInput)
	if err != nil {
		t.Errorf("failed to read input: %v", err)
	}

	got := accessibleToForklift(input)
	if !utils.EqualSlices(expected, got) {
		t.Errorf("accessibleToForklift\nwant %v,\n got %v", expected, got)
	}
}

func TestRemoveAll(t *testing.T) {
	input, err := utils.ReadInput(strings.NewReader(testInput), mapInput)
	if err != nil {
		t.Errorf("failed to read input: %v", err)
	}

	got := removeAll(input)
	if got != 43 {
		t.Errorf("removeAll want 43, got %d", got)
	}
}

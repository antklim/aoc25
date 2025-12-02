package utils_test

import (
	"testing"

	"github.com/antklim/aoc25/internal/utils"
)

func TestEqualSlices(t *testing.T) {
	testCases := []struct {
		a        []int
		b        []int
		expected bool
	}{
		{a: []int{1}, b: []int{1}, expected: true},
		{a: []int{1}, b: []int{2}, expected: false},
		{a: []int{1, 2}, b: []int{1, 2}, expected: true},
		{a: []int{1, 2}, b: []int{2, 2}, expected: false},
	}

	for _, tc := range testCases {
		got := utils.EqualSlices(tc.a, tc.b)
		if got != tc.expected {
			t.Errorf("Equal(%v, %v) want %t, got %t", tc.a, tc.b, tc.expected, got)
		}
	}
}

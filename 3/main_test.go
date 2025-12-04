package main

import "testing"

func TestMaxJoltage(t *testing.T) {
	testCases := []struct {
		a        []int
		expected int
	}{
		{
			a:        []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			expected: 98,
		},
		{
			a:        []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			expected: 89,
		},
		{
			a:        []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			expected: 78,
		},
		{
			a:        []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			expected: 92,
		},
	}

	for _, tc := range testCases {
		got := maxJoltage(tc.a)
		if got != tc.expected {
			t.Errorf("maxJoltage(%v) want %d, got %d", tc.a, tc.expected, got)
		}
	}
}

func TestMaxJoltage12(t *testing.T) {
	testCases := []struct {
		a        []int
		expected int
	}{
		{
			a:        []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			expected: 987654321111,
		},
		{
			a:        []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			expected: 811111111119,
		},
		{
			a:        []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			expected: 434234234278,
		},
		{
			a:        []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			expected: 888911112111,
		},
	}

	for _, tc := range testCases {
		got := maxJoltage12(tc.a)
		if got != tc.expected {
			t.Errorf("maxJoltage12(%v) want %d, got %d", tc.a, tc.expected, got)
		}
	}
}

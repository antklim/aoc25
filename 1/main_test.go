package main

import "testing"

func TestProcessRotation(t *testing.T) {
	t.Run("over min", func(t *testing.T) {
		testCases := []struct {
			rotations []int
			expected  int
		}{
			{rotations: []int{-48, -1}, expected: 0},
			{rotations: []int{-48, -2}, expected: 1},
			{rotations: []int{-48, -3}, expected: 0},
			{rotations: []int{-50}, expected: 1},
		}

		for _, tc := range testCases {
			got := processRotations(tc.rotations)
			if got != tc.expected {
				t.Errorf("processRotations(%v): want %d, got %d", tc.rotations, tc.expected, got)
			}
		}
	})

	t.Run("over max", func(t *testing.T) {
		testCases := []struct {
			rotations []int
			expected  int
		}{
			{rotations: []int{48, 1}, expected: 0},
			{rotations: []int{48, 2}, expected: 1},
			{rotations: []int{48, 3}, expected: 0},
			{rotations: []int{50}, expected: 1},
		}

		for _, tc := range testCases {
			got := processRotations(tc.rotations)
			if got != tc.expected {
				t.Errorf("processRotations(%v): want %d, got %d", tc.rotations, tc.expected, got)
			}
		}
	})

	t.Run("golden", func(t *testing.T) {
		rotations := []int{
			-68,
			-30,
			48,
			-5,
			60,
			-55,
			-1,
			-99,
			14,
			-82,
		}
		got := processRotations(rotations)
		if got != 3 {
			t.Errorf("processRotations(%v): want 3, got %d", rotations, got)
		}
	})
}

func TestNextPos(t *testing.T) {
	testCases := []struct {
		pos      int
		step     int
		expected int
	}{
		{pos: 50, step: 1, expected: 51},
		{pos: 50, step: 101, expected: 51},
		{pos: 50, step: -1, expected: 49},
		{pos: 50, step: -101, expected: 49},

		{pos: 50, step: -68, expected: 82},
		{pos: 82, step: -30, expected: 52},
		{pos: 52, step: 48, expected: 0},
		{pos: 0, step: -5, expected: 95},
		{pos: 95, step: 60, expected: 55},
		{pos: 55, step: -55, expected: 0},
		{pos: 0, step: -1, expected: 99},
		{pos: 99, step: -99, expected: 0},
		{pos: 0, step: 14, expected: 14},
		{pos: 14, step: -82, expected: 32},
	}

	for _, tc := range testCases {
		got := nextPos(tc.pos, tc.step)
		if got != tc.expected {
			t.Errorf("nextPos(%d, %d): want %d, got %d", tc.pos, tc.step, tc.expected, got)
		}
	}
}

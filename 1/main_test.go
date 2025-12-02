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
			got, _ := processRotations(tc.rotations)
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
			got, _ := processRotations(tc.rotations)
			if got != tc.expected {
				t.Errorf("processRotations(%v): want %d, got %d", tc.rotations, tc.expected, got)
			}
		}
	})

	t.Run("nextPos golden", func(t *testing.T) {
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
		got, crossedZeroTimes := processRotations(rotations)
		if got != 3 {
			t.Errorf("processRotations(%v): want 3, got %d", rotations, got)
		}
		if crossedZeroTimes != 3 {
			t.Errorf("processRotations(%v): want crossed zero times 3, got %d", rotations, crossedZeroTimes)
		}
	})
}

func TestNextPos(t *testing.T) {
	testCases := []struct {
		pos         int
		step        int
		expected    int
		crossedZero int
	}{
		{pos: 50, step: 1, expected: 51, crossedZero: 0},
		{pos: 50, step: 101, expected: 51, crossedZero: 1},
		{pos: 50, step: -1, expected: 49, crossedZero: 0},
		{pos: 50, step: -101, expected: 49, crossedZero: 1},

		{pos: 50, step: -68, expected: 82, crossedZero: 1},
		{pos: 82, step: -30, expected: 52, crossedZero: 0},
		{pos: 52, step: 48, expected: 0, crossedZero: 0},
		{pos: 0, step: -5, expected: 95, crossedZero: 0},
		{pos: 95, step: 60, expected: 55, crossedZero: 1},
		{pos: 55, step: -55, expected: 0, crossedZero: 0},
		{pos: 0, step: -1, expected: 99, crossedZero: 0},
		{pos: 99, step: -99, expected: 0, crossedZero: 0},
		{pos: 0, step: 14, expected: 14, crossedZero: 0},
		{pos: 14, step: -82, expected: 32, crossedZero: 1},
	}

	for _, tc := range testCases {
		got := nextPos(tc.pos, tc.step)
		if got != tc.expected {
			t.Errorf("nextPos(%d, %d): want next position %d, got %d", tc.pos, tc.step, tc.expected, got)
		}
	}
}

func TestCrossedZeroTimes(t *testing.T) {
	testCases := []struct {
		pos      int
		step     int
		expected int
	}{
		{pos: 50, step: 1, expected: 0},
		{pos: 50, step: 101, expected: 1},
		{pos: 50, step: -1, expected: 0},
		{pos: 50, step: -101, expected: 1},

		{pos: 50, step: -68, expected: 1},
		{pos: 82, step: -30, expected: 0},
		{pos: 52, step: 48, expected: 0},
		{pos: 0, step: -5, expected: 0},
		{pos: 95, step: 60, expected: 1},
		{pos: 55, step: -55, expected: 0},
		{pos: 0, step: -1, expected: 0},
		{pos: 99, step: -99, expected: 0},
		{pos: 0, step: 14, expected: 0},
		{pos: 14, step: -82, expected: 1},
	}

	for _, tc := range testCases {
		got := crossesZeroTimes(tc.pos, tc.step)
		if got != tc.expected {
			t.Errorf("crossedZeroTimes(%d, %d): want crossed zero times %d, got %d", tc.pos, tc.step, tc.expected, got)
		}
	}
}

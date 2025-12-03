package main

import (
	"sort"
	"testing"

	"github.com/antklim/aoc25/internal/utils"
)

// 11-22
// 95-115
// 998-1012
// 1188511880-1188511890
// 222220-222224
// 1698522-1698528
// 446443-446449
// 38593856-38593862
// 565653-565659
// 824824821-824824827
// 2121212118-2121212124

func TestIsValidID(t *testing.T) {
	testCases := []struct {
		id       string
		expected bool
	}{
		{id: "1", expected: true},
		{id: "101", expected: true},
		{id: "11", expected: false},
		{id: "1122", expected: true},
		{id: "1010", expected: false},
		{id: "222220", expected: true},
		{id: "1188511885", expected: false},
	}

	for _, tc := range testCases {
		got := isValidID(tc.id)
		if got != tc.expected {
			t.Errorf("isValid(%s) want %t, got %t", tc.id, tc.expected, got)
		}
	}
}

func TestIDRangeInvalidIDs(t *testing.T) {
	testCases := []struct {
		idRange  string
		expected []int
	}{
		{idRange: "11-22", expected: []int{11, 22}},
		{idRange: "95-115", expected: []int{99}},
		{idRange: "998-1012", expected: []int{1010}},
		{idRange: "1188511880-1188511890", expected: []int{1188511885}},
		{idRange: "222220-222224", expected: []int{222222}},
		{idRange: "1698522-1698528", expected: nil},
		{idRange: "446443-446449", expected: []int{446446}},
		{idRange: "38593856-38593862", expected: []int{38593859}},
		{idRange: "565653-565659", expected: nil},
		{idRange: "824824821-824824827", expected: nil},
		{idRange: "2121212118-2121212124", expected: nil},
	}

	for _, tc := range testCases {
		idRange, err := NewIDRange(tc.idRange)
		if err != nil {
			t.Errorf("NewIDRange(%s) failed %v", tc.idRange, err)
		}

		got := idRange.InvalidIDs()
		sort.Ints(got)
		sort.Ints(tc.expected)

		if !utils.EqualSlices(got, tc.expected) {
			t.Errorf("IDRange(%s).InvalidIDs want %v, got %v", tc.idRange, tc.expected, got)
		}
	}
}

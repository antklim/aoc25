package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/antklim/aoc25/internal/utils"
)

func main() {
	ranges, err := utils.ReadInput("input.txt", mapIDs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v", err)
		os.Exit(1)
	}

	sum := 0
	for _, id := range invalidIDs(ranges[0]) {
		sum += id
	}

	fmt.Printf("sum of invalid IDs %d\n", sum)
}

func mapIDs(s string) ([]IDRange, error) {
	v := strings.Split(s, ",")
	result := make([]IDRange, 0, len(v))
	for _, idRange := range v {
		r, err := NewIDRange(idRange)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}

func invalidIDs(ranges []IDRange) []int {
	var result []int

	for _, r := range ranges {
		if v := r.InvalidIDs(); len(v) > 0 {
			result = append(result, v...)
		}
	}

	return result
}

type IDRange struct {
	start int
	end   int
}

func NewIDRange(s string) (IDRange, error) {
	v := strings.Split(s, "-")
	if len(v) != 2 {
		return IDRange{}, errors.New("invalid format of id range")
	}

	start, err := strconv.Atoi(v[0])
	if err != nil {
		return IDRange{}, fmt.Errorf("failed to parse start of the id range: %w", err)
	}
	end, err := strconv.Atoi(v[1])
	if err != nil {
		return IDRange{}, fmt.Errorf("failed to parse end of the id range: %w", err)
	}

	return IDRange{start, end}, nil
}

func (r IDRange) InvalidIDs() []int {
	var result []int

	for i := r.start; i <= r.end; i++ {
		if !isValidID(strconv.FormatInt(int64(i), 10)) {
			result = append(result, i)
		}
	}

	return result
}

func isValidID(id string) bool {
	// if len(id)%2 != 0 {
	// 	return true
	// }
	// middle := len(id) / 2
	// return id[:middle] != id[middle:]

	for i := range len(id) / 2 {
		v := strings.ReplaceAll(id, id[0:i+1], "")
		if len(v) == 0 {
			return false
		}
	}

	return true
}

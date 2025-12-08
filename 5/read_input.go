package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readInputFile(file string) ([]Range, []int, error) {
	fi, err := os.Lstat(file)
	if err != nil {
		return nil, nil, err
	}
	if m := fi.Mode(); !m.IsRegular() {
		return nil, nil, fmt.Errorf("want regular file, got: %v", m)
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	return readInput(f)
}

func readInput(r io.Reader) ([]Range, []int, error) {
	var freshIngredientRanges []Range
	var ingredientIDs []int
	state := readingRanges

	br := bufio.NewReader(r)
	for i := 0; ; i++ {
		l, _, err := br.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, nil, err
		}

		s := string(l)
		if s == "" {
			if state == readingRanges {
				state = readingIngredientIDs
				continue
			} else {
				return nil, nil, errors.New("invalid input file format")
			}
		}

		switch state {
		case readingRanges:
			v, err := mapFreshRange(s)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to map file string #%d to fresh ingredient range: %w", i, err)
			}
			freshIngredientRanges = append(freshIngredientRanges, v)
		case readingIngredientIDs:
			v, err := mapIngredientID(s)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to map file string #%d to ingredient id: %w", i, err)
			}
			ingredientIDs = append(ingredientIDs, v)
		default:
			panic("unreachable")
		}
	}

	return freshIngredientRanges, ingredientIDs, nil
}

func mapFreshRange(s string) ([2]int, error) {
	v := strings.Split(s, "-")
	if len(v) != 2 {
		return [2]int{}, errors.New("invalid fresh ingredient range format")
	}
	from, err := strconv.Atoi(v[0])
	if err != nil {
		return [2]int{}, err
	}
	to, err := strconv.Atoi(v[1])
	if err != nil {
		return [2]int{}, err
	}
	return [2]int{from, to}, nil
}

func mapIngredientID(s string) (int, error) {
	return strconv.Atoi(s)
}

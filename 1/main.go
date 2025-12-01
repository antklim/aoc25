package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	MIN = 0
	MAX = 99
	POS = 50
)

func main() {
	rotations, err := readInput("input.txt", mapRotation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v", err)
		os.Exit(1)
	}

	pwd := processRotations(rotations)
	fmt.Printf("password: %d\n", pwd)
	os.Exit(0)
}

// readInput reads file containing input information and parses each line of file to T.
func readInput[R any](file string, mapFunc func(string) (R, error)) ([]R, error) {
	fi, err := os.Lstat(file)
	if err != nil {
		return nil, err
	}
	if m := fi.Mode(); !m.IsRegular() {
		return nil, fmt.Errorf("want regular file, got: %v", m)
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readFile(f, mapFunc)
}

func readFile[R any](r io.Reader, mapFunc func(string) (R, error)) ([]R, error) {
	var result []R
	br := bufio.NewReader(r)
	for i := 0; ; i++ {
		l, _, err := br.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		r, err := mapFunc(string(l))
		if err != nil {
			return nil, fmt.Errorf("failed to map file string #%d to result type: %w", i, err)
		}
		result = append(result, r)
	}

	return result, nil
}

func mapRotation(s string) (int, error) {
	if s[0] != 'R' && s[0] != 'L' {
		return 0, fmt.Errorf("unknown rotation direction %c", s[0])
	}
	sign := 1
	if s[0] == 'L' {
		sign = -1
	}

	i, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, err
	}
	return i * sign, nil
}

func processRotations(a []int) int {
	result := 0
	pos := POS
	if pos == 0 {
		result++
	}

	for _, v := range a {
		pos = nextPos(pos, v)
		if pos == 0 {
			result++
		}
		// fmt.Printf("step %d: v %d, pos %d\n", i, v, pos)
	}

	return result
}

func nextPos(p, s int) int {
	s = s % 100
	if p+s > MAX {
		return p + s - MAX - 1
	}
	if p+s < MIN {
		return MAX + p + s + 1
	}
	return p + s
}

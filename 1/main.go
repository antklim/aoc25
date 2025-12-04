package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/antklim/aoc25/internal/utils"
)

const (
	MIN = 0
	MAX = 99
	POS = 50
)

func main() {
	rotations, err := utils.ReadInputFile("input.txt", mapRotation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v", err)
		os.Exit(1)
	}

	pwd, crossedZero := processRotations(rotations)
	fmt.Printf("password: %d, crossed zero: %d, total: %d\n", pwd, crossedZero, pwd+crossedZero)
	os.Exit(0)
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

func processRotations(a []int) (int, int) {
	result, crossedZero := 0, 0
	pos := POS
	if pos == 0 {
		result++
	}

	for _, v := range a {
		crossedZero += crossesZeroTimes(pos, v)

		pos = nextPos(pos, v)
		if pos == 0 {
			result++
		}
		// fmt.Printf("step %d: v %d, pos %d\n", i, v, pos)
	}

	return result, crossedZero
}

func nextPos(pos, step int) int {
	s := step % 100

	if pos+s > MAX {
		return pos + s - MAX - 1

	}
	if pos+s < MIN {
		return MAX + pos + s + 1
	}

	return pos + s
}

func crossesZeroTimes(pos, step int) int {
	c := abs(step) / 100
	s := step % 100
	n := nextPos(pos, step)

	if (pos+s > MAX || pos+s < MIN) && n != 0 && pos != 0 {
		c++
	}

	return c
}

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return v * -1
}

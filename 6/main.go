package main

import (
	"fmt"
	"os"
)

func main() {
	expressions, err := readInputFile("input.txt", 4)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v", err)
		os.Exit(1)
	}

	sum := sumExpressions(expressions)
	fmt.Printf("sum of expressions: %d\n", sum)
	os.Exit(0)
}

type Expression struct {
	operands  []uint
	operation rune
}

func (expr Expression) Result() uint {
	switch expr.operation {
	case '+':
		result := uint(0)
		for _, x := range expr.operands {
			result += x
		}
		return result
	case '*':
		result := uint(1)
		for _, x := range expr.operands {
			result *= x
		}
		return result
	default:
		panic("unreachable")
	}
}

func sumExpressions(a []Expression) uint {
	result := uint(0)
	for _, expression := range a {
		result += expression.Result()
	}
	return result
}

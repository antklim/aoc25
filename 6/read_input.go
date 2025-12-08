package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

const (
	readingOperands = iota
	readingOperation
)

func readInputFile(file string, opsLine int) ([]Expression, error) {
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

	return readInput(f, opsLine)
}

func readInput(r io.Reader, opsLine int) ([]Expression, error) {
	var operands [][]uint
	var operations []rune

	state := readingOperands
	br := bufio.NewReader(r)
	for i := 0; ; i++ {
		if i == opsLine {
			state = readingOperation
		}

		l, _, err := br.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		switch state {
		case readingOperands:
			v, err := readOperands(l)
			if err != nil {
				return nil, fmt.Errorf("invalid operand at line %d", i)
			}
			operands = append(operands, v)
		case readingOperation:
			v, err := readOperation(l)
			if err != nil {
				return nil, fmt.Errorf("invalid operation at line %d", i)
			}
			operations = v
		default:
			panic("unreachable")
		}
	}

	return mapToExpressions(operands, operations)
}

func readOperands(a []byte) ([]uint, error) {
	var result []uint
	var operand string
	for _, c := range a {
		if c == ' ' && operand != "" {
			v, err := strconv.Atoi(operand)
			if err != nil {
				return nil, err
			}
			result = append(result, uint(v))
			operand = ""
			continue
		}

		if unicode.IsDigit(rune(c)) {
			operand = fmt.Sprintf("%s%c", operand, c)
		}
	}

	if operand != "" {
		v, err := strconv.Atoi(operand)
		if err != nil {
			return nil, err
		}
		result = append(result, uint(v))
	}

	return result, nil
}

// func readOperands2(a []byte) ([]uint, error) {
// 	var result []uint
// 	for x := range strings.SplitSeq(string(a), " ") {
// 		if strings.TrimSpace(x) == "" {
// 			continue
// 		}
// 		v, err := strconv.Atoi(x)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, uint(v))
// 	}
// 	return result, nil
// }

func readOperation(a []byte) ([]rune, error) {
	var result []rune
	for _, c := range a {
		if c == '*' || c == '+' {
			result = append(result, rune(c))
		}
	}
	return result, nil
}

func mapToExpressions(operands [][]uint, operations []rune) ([]Expression, error) {
	var result []Expression

	for i, v := range operands {
		if len(v) != len(operations) {
			return nil, fmt.Errorf("number of operands does not match number of operations on line %d: want %d, got %d", i, len(operations), len(v))
		}
	}

	for i, operation := range operations {
		expressionOperands := make([]uint, 0, len(operands))
		for j := range len(operands) {
			expressionOperands = append(expressionOperands, operands[j][i])
		}
		result = append(result, Expression{
			operands:  expressionOperands,
			operation: operation,
		})
	}

	return result, nil
}

package main

import (
	"strings"
	"testing"

	"github.com/antklim/aoc25/internal/utils"
)

const testInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func TestReadInput(t *testing.T) {
	expected := []Expression{
		{operands: []uint{123, 45, 6}, operation: '*'},
		{operands: []uint{328, 64, 98}, operation: '+'},
		{operands: []uint{51, 387, 215}, operation: '*'},
		{operands: []uint{64, 23, 314}, operation: '+'},
	}

	got, err := readInput(strings.NewReader(testInput), 3)
	if err != nil {
		t.Errorf("failed to read input: %v", err)
	}
	if len(got) != len(expected) {
		t.Errorf("invalid amount of expressions want %d, got %d", len(expected), len(got))
	}
	for i, wantExpression := range expected {
		if !utils.EqualSlices(wantExpression.operands, got[i].operands) {
			t.Errorf("invalid expression operands at line %d want %v,\n got %v", i, wantExpression.operands, got[i].operands)
		}
		if wantExpression.operation != got[i].operation {
			t.Errorf("invalid expression operation at line %d want %c, got %c", i, wantExpression.operation, got[i].operation)
		}
	}
}

func TestSumExpressions(t *testing.T) {
	expected := uint(4277556)
	expressions, err := readInput(strings.NewReader(testInput), 3)
	if err != nil {
		t.Errorf("failed to read input: %v", err)
	}
	got := sumExpressions(expressions)
	if got != expected {
		t.Errorf("invalid sum of expressions want %d, got %d", expected, got)
	}
}

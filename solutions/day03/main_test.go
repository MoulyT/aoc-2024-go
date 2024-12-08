package main

import (
	"testing"
)

const testInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestGetSumOfMul(t *testing.T) {
	expected := 2*4 + 5*5 + 11*8 + 8*5

	got, err := getSumOfMul(testInput)
	if err != nil {
		t.Errorf("getSumOfMul() returned unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("getSumOfMul() = %v, want %v", got, expected)
	}
}

const testInput2 = `
xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
`

func TestGetSumOfMul2(t *testing.T) {
	expected := 2*4 + 8*5

	got, err := getSumOfMul2(testInput2)
	if err != nil {
		t.Errorf("getSumOfMul2() returned unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("getSumOfMul2() = %v, want %v", got, expected)
	}
}

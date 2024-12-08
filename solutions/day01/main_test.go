package main

import (
	"reflect"
	"testing"
)

const testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestGetLists(t *testing.T) {
	expected1 := []int{3, 4, 2, 1, 3, 3}
	expected2 := []int{4, 3, 5, 3, 9, 3}
	col1, col2, err := getLists(testInput)
	if err != nil {
		t.Error("getLists() error", err)
	}
	if !reflect.DeepEqual(col1, expected1) {
		t.Errorf("getLists() col1 = %v, want %v", col1, expected1)
	}
	if !reflect.DeepEqual(col2, expected2) {
		t.Errorf("getLists() col2 = %v, want %v", col2, expected2)
	}
}

func TestTotalDistance(t *testing.T) {
	col1 := []int{3, 4, 2, 1, 3, 3}
	col2 := []int{4, 3, 5, 3, 9, 3}
	expected := 11
	if total := totalDistance(col1, col2); total != expected {
		t.Errorf("totalDistance() = %d, want %d", total, expected)
	}
}

func TestCalculateRepetitions(t *testing.T) {
	col := []int{4, 3, 5, 3, 9, 3}
	testCases := []struct {
		n        int
		expected int
	}{
		{3, 3},
		{4, 1},
		{9, 1},
		{5, 1},
		{7, 0},
	}

	for _, tc := range testCases {
		if got := calculateRepetitions(tc.n, col); got != tc.expected {
			t.Errorf("calculateRepetitions(%d) = %d, want %d", tc.n, got, tc.expected)
		}
	}
}

func TestCalculateSimilarity(t *testing.T) {
	col1 := []int{3, 4, 2, 1, 3, 3}
	col2 := []int{4, 3, 5, 3, 9, 3}
	expected := 31

	if got := calculateSimilarity(col1, col2); got != expected {
		t.Errorf("calculateSimilarity() = %d, want %d", got, expected)
	}
}

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
	expected := &NumberPairs{
		left:  []int{3, 4, 2, 1, 3, 3},
		right: []int{4, 3, 5, 3, 9, 3},
	}

	pairs, err := getLists(testInput)
	if err != nil {
		t.Error("getLists() error", err)
	}

	if !reflect.DeepEqual(pairs, expected) {
		t.Errorf("getLists() = %v, want %v", pairs, expected)
	}
}

func TestTotalDistance(t *testing.T) {
	pairs := &NumberPairs{
		left:  []int{3, 4, 2, 1, 3, 3},
		right: []int{4, 3, 5, 3, 9, 3},
	}
	expected := 11

	if total := totalDistance(pairs); total != expected {
		t.Errorf("totalDistance() = %d, want %d", total, expected)
	}
}

func TestCalculateSimilarity(t *testing.T) {
	pairs := &NumberPairs{
		left:  []int{3, 4, 2, 1, 3, 3},
		right: []int{4, 3, 5, 3, 9, 3},
	}
	expected := 31

	if got := calculateSimilarity(pairs); got != expected {
		t.Errorf("calculateSimilarity() = %d, want %d", got, expected)
	}
}

package main

import (
	"aoc2024/internal/aoc"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type NumberPairs struct {
	left  []int
	right []int
}

func getLists(s string) (*NumberPairs, error) {
	lines := strings.Split(s, "\n")
	// Prealoca los slices basándose en el número de líneas
	pairs := &NumberPairs{
		left:  make([]int, 0, len(lines)),
		right: make([]int, 0, len(lines)),
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}

		n1, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing first number: %w", err)
		}
		n2, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing second number: %w", err)
		}

		pairs.left = append(pairs.left, n1)
		pairs.right = append(pairs.right, n2)
	}
	return pairs, nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func totalDistance(pairs *NumberPairs) int {
	sort.Ints(pairs.left)
	sort.Ints(pairs.right)

	total := 0
	for i := 0; i < len(pairs.left); i++ {
		total += abs(pairs.left[i] - pairs.right[i])
	}
	return total
}

func calculateSimilarity(pairs *NumberPairs) int {
	freqMap := make(map[int]int, len(pairs.right))
	for _, num := range pairs.right {
		freqMap[num]++
	}

	similarity := 0
	for _, num := range pairs.left {
		if freq, exists := freqMap[num]; exists {
			similarity += num * freq
		}
	}
	return similarity
}

func main() {
	starts1 := time.Now()
	input, err := aoc.ReadInput(1)
	if err != nil {
		panic(err)
	}
	numberPairs, err := getLists(input)
	if err != nil {
		panic(err)
	}

	totalDistance := totalDistance(numberPairs)
	duration1 := time.Since(starts1)
	fmt.Println("Part1: ", totalDistance, "time: ", duration1)

	starts2 := time.Now()
	input2, err := aoc.ReadInput(1)
	if err != nil {
		panic(err)
	}
	numberPairs2, err := getLists(input2)
	if err != nil {
		panic(err)
	}
	similarity := calculateSimilarity(numberPairs2)
	duration2 := time.Since(starts2)
	fmt.Println("Part2: ", similarity, "time: ", duration2)
}

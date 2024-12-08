package main

import (
	"aoc2024/internal/aoc"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getSumOfMul(s string) (int, error) {
	sumOfMul := 0
	mulThreeDigitsPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := mulThreeDigitsPattern.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, fmt.Errorf("error parsing number: %w", err)
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, fmt.Errorf("error parsing number: %w", err)
		}
		sumOfMul += num1 * num2

	}
	return sumOfMul, nil
}

func getSumOfMul2(s string) (int, error) {
	sections := strings.Split(s, "don't()")
	sumOfMul := 0
	isOn := true
	for _, section := range sections {
		subsections := strings.Split(section, "do()")

		if isOn {
			sum, err := getSumOfMul(subsections[0])
			if err != nil {
				return 0, fmt.Errorf("error parsing number: %w", err)
			}
			sumOfMul += sum
		}

		for _, subsection := range subsections[1:] {
			sum, err := getSumOfMul(subsection)
			if err != nil {
				return 0, fmt.Errorf("error parsing number: %w", err)
			}
			sumOfMul += sum
		}

		isOn = false
	}

	return sumOfMul, nil
}

func main() {
	starts := time.Now()
	input, err := aoc.ReadInput(3)
	if err != nil {
		panic(err)
	}
	sumOfMul, err := getSumOfMul(input)
	if err != nil {
		panic(err)
	}
	duration := time.Since(starts)
	fmt.Println("Part 1:", sumOfMul, "Time:", duration)

	starts2 := time.Now()
	input2, err := aoc.ReadInput(3)
	if err != nil {
		panic(err)
	}
	sumOfMul2, err := getSumOfMul2(input2)
	if err != nil {
		panic(err)
	}
	duration2 := time.Since(starts2)
	fmt.Println("Part 2:", sumOfMul2, "Time:", duration2)
}

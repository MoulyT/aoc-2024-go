package main

import (
	"aoc2024/internal/aoc"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getReports(s string) ([][]int, error) {
	s = strings.TrimSuffix(s, "\n")
	lines := strings.Split(s, "\n")
	reports := make([][]int, 0, len(lines))

	for i, line := range lines {
		numString := strings.Fields(line)
		report := make([]int, 0, len(numString))

		for _, numChar := range numString {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, fmt.Errorf("error parsing number: %w", err)
			}
			report = append(report, num)
		}
		if len(report) == 0 {
			fmt.Printf("Report vacío generado en línea %d: '%s'\n", i+1, line)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func checkReportSafety(report []int) (bool, error) {
	if len(report) == 0 {
		return false, fmt.Errorf("empty report: %v", report)
	}
	if len(report) == 1 {
		return true, nil
	}
	minLevelDiff := 1
	maxLevelDiff := 3

	isAscending := true
	isDescending := true

	for i, num := range report[:len(report)-1] {
		if abs(num-report[i+1]) < minLevelDiff || abs(num-report[i+1]) > maxLevelDiff {
			return false, nil
		}

		if num > report[i+1] {
			isAscending = false
		}

		if num < report[i+1] {
			isDescending = false
		}

		if !isAscending && !isDescending {
			return false, nil
		}

	}

	return true, nil
}

func checkReportsSafety(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		value, err := checkReportSafety(report)
		if err != nil {
			fmt.Println("error checking report safety")
			panic(err)
		}
		if value {
			safeReports++
		}
	}
	return safeReports
}

func checkReportsSafety2(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		value, err := checkReportSafety(report)
		if err != nil {
			fmt.Println("error checking report safety")
			panic(err)
		}
		if value {
			safeReports++
		} else {
			for i := range report {
				processedReport := make([]int, 0, len(report)-1)
				processedReport = append(processedReport, report[:i]...)
				processedReport = append(processedReport, report[i+1:]...)
				newValue, err := checkReportSafety(processedReport)
				if err != nil {
					fmt.Println("error checking report safety")
					panic(err)
				}
				if newValue {
					safeReports++
					break
				}
			}
		}
	}
	return safeReports
}

func main() {
	starts := time.Now()
	input, err := aoc.ReadInput(2)
	if err != nil {
		panic(err)
	}

	reports, err := getReports(input)
	if err != nil {
		fmt.Println("error getting reports")
		panic(err)
	}
	numSafeReports := checkReportsSafety(reports)
	duration1 := time.Since(starts)
	fmt.Println("Part1: ", numSafeReports, "time: ", duration1)

	starts2 := time.Now()
	input2, err := aoc.ReadInput(2)
	if err != nil {
		panic(err)
	}
	reports2, err := getReports(input2)
	if err != nil {
		fmt.Println("error getting reports")
		panic(err)
	}

	numSafeReports2 := checkReportsSafety2(reports2)
	duration2 := time.Since(starts2)
	fmt.Println("Part2: ", numSafeReports2, "time: ", duration2)
}

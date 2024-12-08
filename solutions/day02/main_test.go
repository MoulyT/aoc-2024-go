package main

import (
	"reflect"
	"testing"
)

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestGetReports(t *testing.T) {
	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	got, err := getReports(testInput)
	if err != nil {
		t.Errorf("getReports() returned unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("getReports() = %v, want %v", got, expected)
	}
}

func TestCheckReportSafety(t *testing.T) {
	tests := []struct {
		name   string
		report []int
		want   bool
	}{
		{
			name:   "Test from input - first row",
			report: []int{7, 6, 4, 2, 1},
			want:   true,
		},
		{
			name:   "Test from input - second row",
			report: []int{1, 2, 7, 8, 9},
			want:   false,
		},
		{
			name:   "Valid report with differences between 1-3",
			report: []int{9, 7, 6, 2, 1},
			want:   false,
		},
		{
			name:   "Invalid report - it's increasing and decreasing",
			report: []int{1, 3, 2, 4, 5},
			want:   false,
		},
		{
			name:   "Invalid report - difference too large",
			report: []int{8, 6, 4, 4, 1},
			want:   false,
		},
		{
			name:   "Valid report with differences between 1-3",
			report: []int{1, 3, 6, 7, 9},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := checkReportSafety(tt.report)
			if got != tt.want {
				t.Errorf("checkReportSafety(%v) = %v, want %v",
					tt.report, got, tt.want)
			}
		})
	}
}

func TestCheckReportsSafety(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	want := 2
	got := checkReportsSafety(reports)

	if got != want {
		t.Errorf("checkReportsSafety() = %v, want %v", got, want)
	}
}


func TestCheckReportsSafety2(t *testing.T) {
  reports := [][]int{
    {7, 6, 4, 2, 1},
    {1, 2, 7, 8, 9},
    {9, 7, 6, 2, 1},
    {1, 3, 2, 4, 5},
    {8, 6, 4, 4, 1},
    {1, 3, 6, 7, 9},
  }

  want := 4
  got := checkReportsSafety2(reports)

  if got != want {
    t.Errorf("checkReportsSafety2() = %v, want %v", got, want)
  }
}

package aoc

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadInput(day int) (string, error) {
	path := filepath.Join("solutions", fmt.Sprintf("day%02d", day), "input.txt")
	data, err := os.ReadFile(path)
	return string(data), err
}

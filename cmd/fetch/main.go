package main

import (
    "flag"
    "fmt"
    "log"
    "os"

    "aoc2024/internal/aoc"
)

func main() {
    day := flag.Int("day", 0, "Day to fetch (1-25)")
    flag.Parse()

    if *day < 1 || *day > 25 {
        log.Fatal("Day must be between 1 and 25")
    }

    token := os.Getenv("AOC_SESSION")
    if token == "" {
        log.Fatal("AOC_SESSION environment variable not set")
    }

    client := aoc.NewClient(token, 2024)
    input, err := client.FetchInput(*day)
    if err != nil {
        log.Fatalf("Failed to fetch input: %v", err)
    }

    if err := client.SaveInput(*day, input); err != nil {
        log.Fatalf("Failed to save input: %v", err)
    }

    fmt.Printf("Successfully fetched and saved input for day %d\n", *day)
}

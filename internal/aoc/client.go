package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Client struct {
	SessionToken string
	Year         int
}

func NewClient(token string, year int) *Client {
	return &Client{
		SessionToken: token,
		Year:         year,
	}
}

func (c *Client) FetchInput(day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", c.Year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", c.SessionToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func (c *Client) SaveInput(day int, input string) error {
	dir := filepath.Join("solutions", fmt.Sprintf("day%02d", day))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, "input.txt"), []byte(input), 0644)
}

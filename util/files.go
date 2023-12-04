package util

import (
	"bufio"
	"os"
)

// ReadLines opens the file at path and returns a slice of strings, each
// element of which is a line from the file. If an error occurs, it is
// returned.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

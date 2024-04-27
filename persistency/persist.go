package persistency

import (
	"bufio"
	"os"
)

var ReadAllLines = func(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // considering only non-empty lines
			lines = append(lines, line)
		}
	}

	return lines, nil
}

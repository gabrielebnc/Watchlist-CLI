package persistency

import (
	"bufio"
	"os"
)

func ReadAllLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

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

func PersistItem(item string, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(item + "\n")
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

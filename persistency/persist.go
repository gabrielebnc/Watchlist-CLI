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
	_, err = file.Write([]byte(item))
	if err != nil {
		file.Close()
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

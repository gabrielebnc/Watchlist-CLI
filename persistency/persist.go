package persistency

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func addPrefixToPath(prefix string, path string) string {

	//Case where using tilde -> add user home dir in the prefix
    if prefix[:1] == "~" {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            fmt.Println("Error getting user home directory:", err)
            return ""
        }
        prefix = filepath.Join(homeDir, prefix[1:])
    }

    // Join prefix and path
    fullPath := filepath.Join(prefix, path)
    return fullPath
}


func ReadAllLines(path string) ([]string, error) {

	file, err := os.Open(addPrefixToPath("~", path))
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
	file, err := os.OpenFile(addPrefixToPath("~", path), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
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

func RemoveLineAtIndex(path string, index int) (err error) {
	f, err := os.OpenFile(addPrefixToPath("~", path), os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error while removing: file opening failed")
	}
	scanner := bufio.NewScanner(f)
	var bs []byte
	buf := bytes.NewBuffer(bs)

	var lineIndex int
	var text string
	for scanner.Scan() {
		lineIndex++
		text = scanner.Text()
		if lineIndex != index {
			_, err := buf.WriteString(text + "\n")
			if err != nil {
				return fmt.Errorf("error while removing: line replacement")
			}
		}
	}
	f.Truncate(0)
	f.Seek(0, 0)
	buf.WriteTo(f)
	return nil
}

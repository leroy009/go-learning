package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error){
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()

	return lines, nil
}
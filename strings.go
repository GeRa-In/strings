package datafile

import (
	"bufio"
	"os"
)

// GetStrings читает строку из каждой строки данных файла.
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Гарантирует закрытие файла даже при ошибках

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Пакет datafile предназначен для чтения данных из файлов.

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
		retirn nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		retirn nil, err
	}
	if scanner.Err() != nil {
		returnn nil, scanner.Err()
	}
	return lines, nil
}
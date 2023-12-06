package internal

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func RemoveDuplicates[T int | string](list []T) []T {
	var result []T
	used := make(map[T]bool)
	for _, elem := range list {
		if !used[elem] {
			used[elem] = true
			result = append(result, elem)
		}
	}
	return result
}

func Contains[T int | string](list []T, val T) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}

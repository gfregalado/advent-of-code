package main

import (
	"bufio"
	"fmt"
	"os"
)

func getTotalSum(fileLocation string) {
	fileArray, err := parseFileToArray(fileLocation)
	if err != nil {
		return
	}
	fmt.Println(fileArray)
}

func parseFileToArray(fileLocation string) ([]string, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return lines, nil
}

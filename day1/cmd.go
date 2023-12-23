package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fileArray, err := parseFileToArray("1.txt")
	if err != nil {
		return
	}
	fmt.Println(fileArray)
	processedStrings := keepOnlyNumbers(fileArray)
	fmt.Println("Processed strings:", processedStrings)
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

func keepOnlyNumbers(strings []string) []string {
	// Compile a regular expression to match non-numeric characters.
	re, err := regexp.Compile("[^0-9]+")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	var processedStrings []string
	for _, s := range strings {
		// Replace non-numeric characters with an empty string.
		processedString := re.ReplaceAllString(s, "")
		processedStrings = append(processedStrings, processedString)
	}

	return processedStrings
}

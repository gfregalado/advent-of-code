package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseDocumentToTotalNumber(fileLocation string) {
	fileArray, err := parseFileToArray(fileLocation)
	if err != nil {
		return
	}
	fmt.Println(fileArray)
	processedStrings := keepOnlyNumbers(fileArray)
	fmt.Println("Processed strings:", processedStrings)
	parsedNumbers := getParsedNumbersForSum(processedStrings)
	fmt.Println("Parsed numbers:", parsedNumbers)
	totalSum := sumAllEntries(parsedNumbers)
	fmt.Println("Total sum:", totalSum)

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

func keepOnlyNumbers(input []string) []string {
	// Compile a regular expression to match non-numeric characters.
	re, err := regexp.Compile("[^0-9]+")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	var processedStrings []string
	for _, s := range input {
		// Replace non-numeric characters with an empty string.
		processedString := re.ReplaceAllString(s, "")
		processedStrings = append(processedStrings, processedString)
	}

	return processedStrings
}

func getParsedNumbersForSum(input []string) []string {
	var processedNumbers []string
	for _, s := range input {
		switch len(s) {
		case 1:
			// Concatenate the string with itself for length 1.
			temp := s + s
			processedNumbers = append(processedNumbers, temp)

		case 2:
			// Use the string as is for length 2.
			processedNumbers = append(processedNumbers, s)

		default:
			// For longer strings, concatenate the first and last character.
			runes := []rune(s)
			firstRune := string(runes[0])
			lastRune := string(runes[len(runes)-1])
			processedNumbers = append(processedNumbers, firstRune+lastRune)
		}
	}

	return processedNumbers
}

func sumAllEntries(input []string) int {
	var totalSum int
	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			// ... handle error
			panic(err)
		}
		totalSum += i
	}
	return totalSum
}

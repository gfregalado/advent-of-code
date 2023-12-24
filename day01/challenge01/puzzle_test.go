package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedResultFromFile = []string{"five8b", "2733vmmpknvgr", "3oneeighttwo"}
var expectedResultFromRemovingLetters = []string{"8", "2733", "3"}

var expectedResultFromParsedNumForSum = []string{"88", "23", "33"}

func TestTransformFileToArray(t *testing.T) {
	parsedArray, err := parseFileToArray("./test.txt")

	assert.Equal(t, expectedResultFromFile, parsedArray)
	assert.NoError(t, err, "The function should not return an error")
}

func TestRemoveLettersFromArrayElements(t *testing.T) {
	parsedArray := keepOnlyNumbers(expectedResultFromFile)

	assert.Equal(t, parsedArray, expectedResultFromRemovingLetters)
}

func TestKeepOnlyNumbersForSume(t *testing.T) {
	parsedArray := getParsedNumbersForSum(expectedResultFromRemovingLetters)

	assert.Equal(t, parsedArray, expectedResultFromParsedNumForSum)
}

func TestSumAllEntries(t *testing.T) {
	summedArray := sumAllEntries(expectedResultFromParsedNumForSum)

	assert.Equal(t, summedArray, 144)
}

func TestParseDocumentToTotalNumber(t *testing.T) {
	parseDocumentToTotalNumber("./file.txt")
}

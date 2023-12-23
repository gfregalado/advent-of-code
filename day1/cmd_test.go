package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedResult = []string{"five8b", "2733vmmpknvgr", "3oneeighttwo"}

func TestTransformFileToArray(t *testing.T) {
	parsedArray, err := parseFileToArray("./test.txt")

	assert.Equal(t, expectedResult, parsedArray)
	assert.NoError(t, err, "The function should not return an error")
}

func TestRemoveLettersFromArrayElements(t *testing.T) {
	parsedArray := keepOnlyNumbers(expectedResult)

	assert.Equal(t, parsedArray, []string{"8", "2733", "3"})
}

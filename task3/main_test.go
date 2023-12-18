package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ArrayOfLowestAndHighestScores1(t *testing.T) {
	scores := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}

	expectedHighestScores := []int32{4, 21, 36, 42}
	expectedLowestScores := []int32{}

	actual1, actual2 := FindHighestAndLowestScores(scores)

	assert.Equal(t, expectedHighestScores, actual1)
	assert.Equal(t, expectedLowestScores, actual2)
}

func Test_ArrayOfLowestAndHighestScores2(t *testing.T) {
	scores := []int32{12, 24, 10, 24}

	expectedHighestScores := []int32{24}
	expectedLowestScores := []int32{10}

	actual1, actual2 := FindHighestAndLowestScores(scores)

	assert.Equal(t, expectedHighestScores, actual1)
	assert.Equal(t, expectedLowestScores, actual2)
}

func Test_ArrayOfLowestAndHighestScores3(t *testing.T) {
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}

	expectedHighestScores := []int32{20, 25}
	expectedLowestScores := []int32{5, 4, 2, 1}

	actual1, actual2 := FindHighestAndLowestScores(scores)

	assert.Equal(t, expectedHighestScores, actual1)
	assert.Equal(t, expectedLowestScores, actual2)
}

func Test_CountHighestScore(t *testing.T) {
	highestScoreArray := []int32{4, 21, 36, 42}

	expected := int32(4)
	actual := CountHighestScore(highestScoreArray)

	assert.Equal(t, expected, actual)
}

func Test_CountLowestScore(t *testing.T) {
	highestScoreArray := []int32{5, 4, 2, 1}

	expected := int32(4)
	actual := CountHighestScore(highestScoreArray)

	assert.Equal(t, expected, actual)
}

func Test_BreakingScores(t *testing.T) {
	scores := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}

	expected1 := int32(4)
	expected2 := int32(0)

	actual1, actual2 := BreakingRecords(scores)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}

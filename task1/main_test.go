package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AppleCount1(t *testing.T) {
	appleLocation := 5
	apples := []int32{-2, 2, 1}

	expected := []int{3, 7, 6}
	actual := CountAllApples(appleLocation, apples)

	assert.Equal(t, expected, actual)
}

func Test_AppleFilter1(t *testing.T) {
	applesCount := []int{3, 7, 6}
	startingPointHouse := 7
	endingLocationOfHouse := 11

	expected := []int{7}
	actual := FilterApples(applesCount, startingPointHouse, endingLocationOfHouse)

	assert.Equal(t, expected, actual)
}

func Test_CountFilteredApples1(t *testing.T) {
	filteredApples := []int{7}

	expected := 1
	actual := CountFilteredApples(filteredApples)

	assert.Equal(t, expected, actual)
}

func Test_OrangeCount1(t *testing.T) {
	orangeLocation := 15
	oranges := []int32{5, -6}

	expected := []int{20, 9}
	actual := CountAllOranges(orangeLocation, oranges)

	assert.Equal(t, expected, actual)
}

func Test_OrangeFilter1(t *testing.T) {
	orangesCount := []int{20, 9}
	startingPointHouse := 7
	endingLocationOfHouse := 11

	expected := []int{9}
	actual := FilterOranges(orangesCount, startingPointHouse, endingLocationOfHouse)

	assert.Equal(t, expected, actual)
}

func Test_CountFilteredOranges1(t *testing.T) {
	filteredOranges := []int{7, 9, 2}

	expected := 3
	actual := CountFilteredOranges(filteredOranges)

	assert.Equal(t, expected, actual)
}

func Test_CountFilteredOranges2(t *testing.T) {
	filteredOranges := []int{9}

	expected := 1
	actual := CountFilteredOranges(filteredOranges)

	assert.Equal(t, expected, actual)
}

func Test_CountApplesAndOranges(t *testing.T) {
	startingPointOfHouse := 7
	endingLocationOfHouse := 11
	appleLocation := 5
	orangeLocation := 15
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	expected1 := 1
	expected2 := 1

	actual1, actual2 := CounterApplesAndOranges(startingPointOfHouse, endingLocationOfHouse, appleLocation, orangeLocation, apples, oranges)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}

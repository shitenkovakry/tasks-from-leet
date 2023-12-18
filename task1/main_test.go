package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AppleCount1(t *testing.T) {
	appleLocation := 5
	apples := []int{-2, 2, 1}

	expected := []int{3,7,6}
	actual := CountAllApples(appleLocation, apples)

	assert.Equal(t, expected, actual)
}

func Test_OrangeCount1(t *testing.T) {
	orangeLocation := 15
	oranges := []int{5, -6}

	expected := []int{20,9}
	actual := CountAllOranges(orangeLocation, oranges)

	assert.Equal(t, expected, actual)
}

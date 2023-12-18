package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AppleCount(t *testing.T) {
	appleLocation := 5
	apples := []int{-2, 2, 1}

	expected := []int{3,7,6}
	actual := CountAllApples(appleLocation, apples)

	assert.Equal(t, expected, actual)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JumpKangaroo1(t *testing.T) {
	startPositionTheFirst := int32(2)
	jumpDistanceTheFirst := int32(1)
	startPositionTheSecond := int32(1)
	jumpDistanceTheSecond := int32(2)

	expected := "Yes"
	actual := JumbKangaroos(startPositionTheFirst, jumpDistanceTheFirst, startPositionTheSecond, jumpDistanceTheSecond)

	assert.Equal(t, expected, actual)
}

func Test_JumpKangaroo2(t *testing.T) {
	startPositionTheFirst := int32(0)
	jumpDistanceTheFirst := int32(3)
	startPositionTheSecond := int32(4)
	jumpDistanceTheSecond := int32(2)

	expected := "Yes"
	actual := JumbKangaroos(startPositionTheFirst, jumpDistanceTheFirst, startPositionTheSecond, jumpDistanceTheSecond)

	assert.Equal(t, expected, actual)
}

func Test_JumpKangaroo3(t *testing.T) {
	startPositionTheFirst := int32(0)
	jumpDistanceTheFirst := int32(2)
	startPositionTheSecond := int32(5)
	jumpDistanceTheSecond := int32(3)

	expected := "No"
	actual := JumbKangaroos(startPositionTheFirst, jumpDistanceTheFirst, startPositionTheSecond, jumpDistanceTheSecond)

	assert.Equal(t, expected, actual)
}

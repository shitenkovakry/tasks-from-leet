package main

func JumbKangaroos(startPositionTheFirst int32, jumpDistanceTheFirst int32, startPositionTheSecond int32, jumpDistanceTheSecond int32) string {
	meet := "Yes"
	notMeet := "No"

	if startPositionTheFirst == startPositionTheSecond {
		return meet
	}

	sequenceLocationTheFirst := startPositionTheFirst
	sequenceLocationTheSecond := startPositionTheSecond

	deltaFirst := jumpDistanceTheFirst
	deltaSecond := jumpDistanceTheSecond

	if startPositionTheFirst > startPositionTheSecond {
		sequenceLocationTheFirst = startPositionTheSecond
		sequenceLocationTheSecond = startPositionTheFirst

		deltaFirst = jumpDistanceTheSecond
		deltaSecond = jumpDistanceTheFirst
	}

	for {
		sequenceLocationTheFirst += deltaFirst
		sequenceLocationTheSecond += deltaSecond

		if sequenceLocationTheFirst == sequenceLocationTheSecond {
			return meet
		} else if sequenceLocationTheFirst > sequenceLocationTheSecond {
			return notMeet
		}
	}
}

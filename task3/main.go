package main

func BreakingRecords(scores []int32) (int32, int32) {
	highestScoreArray, lowestScoreArray := FindHighestAndLowestScores(scores)

	quantityHighestScores := CountHighestScore(highestScoreArray)
	quantityLowestScores := CountLowestScore(lowestScoreArray)

	return quantityHighestScores, quantityLowestScores
}

func FindHighestAndLowestScores(scores []int32) ([]int32, []int32) {
	highestScoreArray := []int32{}
	lowestScoreArray := []int32{}
	highestScore := scores[0]
	lowestScore := scores[0]

	for _, score := range scores {
		if score > highestScore {
			highestScoreArray = append(highestScoreArray, score)
			highestScore = score

		} else if score < lowestScore {
			lowestScoreArray = append(lowestScoreArray, score)
			lowestScore = score

		}
	}

	return highestScoreArray, lowestScoreArray
}

func CountHighestScore(highestScoreArray []int32) int32 {
	lenOfHighestScoreArray := len(highestScoreArray)

	count := int32(lenOfHighestScoreArray)

	return count
}

func CountLowestScore(lowestScoreArray []int32) int32 {
	lenOfLowestScoreArray := len(lowestScoreArray)

	count := int32(lenOfLowestScoreArray)

	return count
}

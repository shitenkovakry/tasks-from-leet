package main
func CounterApplesAndOranges(startingPointOfHouse int, endingLocationOfHouse int, appleLocation int, orangeLocation int, apples []int, oranges []int) ([]int, []int){
	applesCount := CountAllApples(appleLocation, apples)
	orangesCount := CountAllOranges(orangeLocation, oranges)

	return applesCount, orangesCount
}

func CountAllApples(appleLocation int, apples []int) []int {
	applesCount := []int{}


	for _, apple := range apples {
		resultForApple := appleLocation + apple

		applesCount = append(applesCount, resultForApple)
	}

	return applesCount
}

func CountAllOranges(orangeLocation int, oranges []int) []int {
	orangesCount := []int{}


	for _, orange := range oranges {
		resultForOrange := orangeLocation + orange

		orangesCount = append(orangesCount, resultForOrange)
	}

	return orangesCount
}

func FilterApples(applesCount []int, startingPointHouse int, endingLocationOfHouse int) []int{
	resultApples := []int{}
	for _, apple := range applesCount {
		if apple >= startingPointHouse && apple <= endingLocationOfHouse {
			resultApples = append(resultApples, apple)
		}
	}

	return resultApples
}

func FilterOranges(orangesCount []int, startingPointHouse int, endingLocationOfHouse int) []int{
	resultOranges := []int{}
	for _, orange := range orangesCount {
		if orange >= startingPointHouse && orange <= endingLocationOfHouse {
			resultOranges = append(resultOranges, orange)
		}
	}

	return resultOranges
}

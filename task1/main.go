package main

import "fmt"

func CounterApplesAndOranges(startingPointOfHouse int, endingLocationOfHouse int, appleLocation int, orangeLocation int, apples []int32, oranges []int32) (int, int){
	applesCount := CountAllApples(appleLocation, apples)
	orangesCount := CountAllOranges(orangeLocation, oranges)

	filteredApples := FilterApples(applesCount, startingPointOfHouse, endingLocationOfHouse)
	filteredOranges := FilterOranges(orangesCount, startingPointOfHouse, endingLocationOfHouse)

	countOfFilteredApples := CountFilteredApples(filteredApples)
	countOfFilteredOranges := CountFilteredOranges(filteredOranges)

	return countOfFilteredApples, countOfFilteredOranges
}

func CountAllApples(appleLocation int, apples []int32) []int {
	applesCount := []int{}


	for _, apple := range apples {
		resultForApple := appleLocation + int(apple)

		applesCount = append(applesCount, int(resultForApple))
	}

	return applesCount
}

func CountAllOranges(orangeLocation int, oranges []int32) []int {
	orangesCount := []int{}


	for _, orange := range oranges {
		resultForOrange := orangeLocation + int(orange)

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

func CountFilteredApples(filteredApples []int) int {
	lenOfInputArray := len(filteredApples)

	filtered := lenOfInputArray

	return filtered
}

func CountFilteredOranges(filteredOranges []int) int {
	lenOfInputArray := len(filteredOranges)

	filtered := lenOfInputArray

	return filtered
}

// countApplesAndOranges implements the result for hackerrank task count apples and oranges
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	x, y := CounterApplesAndOranges(int(s), int(t),int(a),int(b),apples, oranges)

	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	countApplesAndOranges(
		7, 11,
		5, 15,
		[]int32{-2, 2, 1},
		[]int32{5, -6},
	)

	/* you expect to see here...
	1
	1
	*/
}

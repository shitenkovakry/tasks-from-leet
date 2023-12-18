package main

/*func CounterApplesAndOranges(startingPointOfHouse int, endingLocationOfHouse int, appleLocation int, orangeLocation int, apples []int, oranges []int) []int{

	//orangeCount := []int{}
} */

func CountAllApples(appleLocation int, apples []int) []int {
	applesCount := []int{}


	for _, apple := range apples {
		resultForApple := appleLocation + apple

		applesCount = append(applesCount, resultForApple)
	}

	return applesCount
}

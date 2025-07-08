package arrays

func Sum(numbers []int) int{
	sum := 0

// _ (blank space because we dont want to get the index.)
// mySlice := []int{1, 2, 3} , myArray := [3]int{1, 2, 3}

	for _, number := range numbers{
		sum += number
	}
	return sum
}
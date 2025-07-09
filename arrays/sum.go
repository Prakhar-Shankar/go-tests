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

func SumAll(numbersToSum ...[]int) []int{
	var sums []int

	for _, numbers := range numbersToSum{
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int{
	var sums []int

	for _, numbers := range numbersToSum{
		if len(numbers) == 0{
			sums = append(sums, 0)
		}else{
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
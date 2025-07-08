package arrays

import "testing"

func TestSum(t *testing.T){
	//[N]type{value1, value2...}
	//[...]type{value1, value2...}
	t.Run("Collection of 5 numbers-array", func(t *testing.T){
		numbers := []int{1,2,3,4,5}
		
		got:= Sum(numbers)
		want := 15

		if got != want{
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size-Slices", func(t *testing.T){
		numbers := []int{1,2,3}

		got := Sum(numbers)
		want := 6

		if got != want{
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}
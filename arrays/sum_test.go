package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		wanted := 6 

		if got != wanted {
			t.Errorf("got %d want %d, given %v", got, wanted, numbers)
		}
	})
}

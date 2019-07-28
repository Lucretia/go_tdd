package arrays

import "testing"

func testSlice(t *testing.T, numbers []int, total int) {
	got := Sum(numbers)
	want := total

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		testSlice(t, numbers, 6)
	})
}

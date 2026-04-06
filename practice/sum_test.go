package main

import "testing"


func TestSums(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		want := 15
		got := Sum(numbers)
		if want != got {
			t.Errorf("want %d got %d given %v", want , got, numbers)
		}
	})

	t.Run("collection on any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		want := 6
		got := Sum(numbers)

		if want != got {
			t.Errorf("want %d got %d given %v", want, got, numbers)
		}
	})
}
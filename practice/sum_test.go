package main

import (
	"slices"
	"testing"
)

func TestSums(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		want := 15
		got := Sum(numbers)
		if want != got {
			t.Errorf("want %d got %d given %v", want, got, numbers)
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

	t.Run("collection of all sum", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("get sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

package main

import (
	"reflect"
	"testing"
)

func TestSums(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d got %d", want, got)
		}
	}
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

		checkSums(t, got, want)
	})

	t.Run("get sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

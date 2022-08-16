package leetcode324

import (
	"fmt"
	"testing"
)

func TestCheckWiggle(t *testing.T) {
	assertBoolean := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("check if wiggle", func(t *testing.T) {
		got := checkWiggle([]int{1, 6, 1, 5, 1, 4})
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("check if not wiggle", func(t *testing.T) {
		got := checkWiggle([]int{1, 5, 1, 1, 6, 4})
		want := false
		assertBoolean(t, got, want)
	})
}

func TestWiggleSort(t *testing.T) {
	assertBoolean := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("1, 5, 1, 1, 6, 4", func(t *testing.T) {
		nums := []int{1, 5, 1, 1, 6, 4}
		wiggleSort(nums)
		got := checkWiggle(nums)
		fmt.Printf("nums: %v\t%t\n\n\n", nums, got)
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("1, 3, 2, 2, 3, 1", func(t *testing.T) {
		nums := []int{1, 3, 2, 2, 3, 1}
		wiggleSort(nums)
		got := checkWiggle(nums)
		fmt.Printf("nums: %v\t%t\n\n\n", nums, got)
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("1, 1, 2, 2, 2, 1", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 2, 1}
		wiggleSort(nums)
		got := checkWiggle(nums)
		fmt.Printf("nums: %v\t%t\n\n\n", nums, got)
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("check if wiggle in place", func(t *testing.T) {
		nums := []int{1, 1, 2, 1, 2, 2, 1}
		wiggleSort(nums)
		got := checkWiggle(nums)
		fmt.Printf("nums: %v\t%t\n\n\n", nums, got)
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("check if wiggle in place", func(t *testing.T) {
		nums := []int{1, 1, 2, 1, 2, 2, 1}
		wiggleSort(nums)
		got := checkWiggle(nums)
		fmt.Printf("nums: %v\t%t\n\n\n", nums, got)
		want := true
		assertBoolean(t, got, want)
	})
}

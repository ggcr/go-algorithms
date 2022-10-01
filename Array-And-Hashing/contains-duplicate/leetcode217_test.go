package leetcode217

import (
	"reflect"
	"testing"
)

func TestCheckWiggle(t *testing.T) {
	assertBoolean := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("BRUTE FORCE (1)", func(t *testing.T) {
		got := containsDuplicateBruteForce([]int{1, 5, 3, 2, 6, 4})
		want := false
		assertBoolean(t, got, want)
	})

	t.Run("MERGE SORT (1)", func(t *testing.T) {
		got, boolean_got := containsDuplicateMergeSort([]int{1, 5, 3, 2, 6, 4})
		want, boolean_want := []int{1, 2, 3, 4, 5, 6}, false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
		assertBoolean(t, boolean_got, boolean_want)
	})

	t.Run("MERGE SORT (2)", func(t *testing.T) {
		got, boolean_got := containsDuplicateMergeSort([]int{1, 2, 3, 1})
		want, boolean_want := []int{1, 1, 2, 3}, true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
		assertBoolean(t, boolean_got, boolean_want)
	})

	t.Run("HASH SET (1)", func(t *testing.T) {
		got := containsDuplicate([]int{1, 6, 1, 5, 1, 4})
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("HASH SET (2)", func(t *testing.T) {
		got := containsDuplicate([]int{1, 5, 3, 2, 6, 4})
		want := false
		assertBoolean(t, got, want)
	})

}

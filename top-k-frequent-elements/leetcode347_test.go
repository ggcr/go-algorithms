package leetcode347

import (
	"reflect"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	assertSlice := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	// t.Run("HASH SET (1)", func(t *testing.T) {
	// 	got := topKFrequentSort([]int{1, 1, 1, 2, 2, 3}, 2)
	// 	want := []int{1, 2}
	// 	assertSlice(t, got, want)
	// })

	// t.Run("HASH SET (2)", func(t *testing.T) {
	// 	got := topKFrequentSort([]int{-1, -1}, 1)
	// 	want := []int{-1}
	// 	assertSlice(t, got, want)
	// })

	// t.Run("HASH SET (3)", func(t *testing.T) {
	// 	got := topKFrequentSort([]int{-1, -4, 2, 1, -1, 2, 0, 4, -4, 2, 1, 4, 2, 3, 4, 6, 8, 9, 5, 6, 77}, 4)
	// 	want := []int{-1}
	// 	assertSlice(t, got, want)
	// })

	t.Run("HASH SET (1)", func(t *testing.T) {
		got := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
		want := []int{1, 2}
		assertSlice(t, got, want)
	})

	t.Run("HASH SET (2)", func(t *testing.T) {
		got := topKFrequent([]int{-1, -1}, 1)
		want := []int{-1}
		assertSlice(t, got, want)
	})

	t.Run("HASH SET (3)", func(t *testing.T) {
		got := topKFrequent([]int{-1, -4, 2, 1, -1, 2, 0, 4, -4, 2, 1, 4, 2, 3, 4, 6, 8, 9, 5, 6, 77}, 4)
		want := []int{-1}
		assertSlice(t, got, want)
	})
}

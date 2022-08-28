package leetcode167

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assertSlice := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run(" (1)", func(t *testing.T) {
		got := twoSum([]int{2, 7, 11, 15}, 9)
		want := []int{1, 2}
		assertSlice(t, got, want)
	})

	t.Run(" (2)", func(t *testing.T) {
		got := twoSum([]int{2, 3, 4}, 6)
		want := []int{1, 3}
		assertSlice(t, got, want)
	})

	t.Run(" (3)", func(t *testing.T) {
		got := twoSum([]int{-1, 0}, -1)
		want := []int{1, 2}
		assertSlice(t, got, want)
	})
}

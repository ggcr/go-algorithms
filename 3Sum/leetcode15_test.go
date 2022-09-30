package leetcode15

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assertSlice := func(t testing.TB, got, want [][]int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := threeSum([]int{-1, 0, 1, 2, -1, -4})
		want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
		assertSlice(t, got, want)
	})

	t.Run("test (2)", func(t *testing.T) {
		got := threeSum([]int{0, 1, 1})
		want := [][]int{}
		assertSlice(t, got, want)
	})

	t.Run("test (3)", func(t *testing.T) {
		got := threeSum([]int{0, 0, 0})
		want := [][]int{{0, 0, 0}}
		assertSlice(t, got, want)
	})
}

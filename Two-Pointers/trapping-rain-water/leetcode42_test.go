package leetcode42

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assert := func(t testing.TB, got, want int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := trap([]int{4, 2, 0, 3, 2, 5})
		want := 9
		assert(t, got, want)
	})

	t.Run("test (2)", func(t *testing.T) {
		got := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		want := 6
		assert(t, got, want)
	})

	t.Run("test (3)", func(t *testing.T) {
		got := trap([]int{4, 2, 3})
		want := 1
		assert(t, got, want)
	})

	t.Run("test (4)", func(t *testing.T) {
		got := trap([]int{5, 4, 1, 2})
		want := 1
		assert(t, got, want)
	})

	t.Run("test (5)", func(t *testing.T) {
		got := trap([]int{9, 8, 2, 6})
		want := 4
		assert(t, got, want)
	})

	t.Run("test (6)", func(t *testing.T) {
		got := trap([]int{0, 7, 1, 4, 6})
		want := 7
		assert(t, got, want)
	})

	t.Run("test (7)", func(t *testing.T) {
		got := trap([]int{2, 6, 3, 8, 2, 7, 2, 5, 0})
		want := 11
		assert(t, got, want)
	})

	t.Run("test (8)", func(t *testing.T) {
		got := trap([]int{2, 9, 6, 3, 6, 7, 6})
		want := 6
		assert(t, got, want)
	})
}

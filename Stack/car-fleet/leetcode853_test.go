package leetcode853

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assert := func(t testing.TB, got, want int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Test", func(t *testing.T) {
		got := carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})
		want := 3
		assert(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := carFleet(100, []int{0, 2, 4}, []int{4, 2, 1})
		want := 1
		assert(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := carFleet(10, []int{3}, []int{3})
		want := 1
		assert(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := carFleet(10, []int{6, 8}, []int{3, 2})
		want := 2
		assert(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := carFleet(10, []int{0, 4, 2}, []int{2, 1, 3})
		want := 1
		assert(t, got, want)
	})
}

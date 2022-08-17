package leetcode217

import (
	"testing"
)

func TestCheckWiggle(t *testing.T) {
	assertBoolean := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		got := containsDuplicate([]int{1, 6, 1, 5, 1, 4})
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("test 2", func(t *testing.T) {
		got := containsDuplicate([]int{1, 5, 3, 2, 6, 4})
		want := false
		assertBoolean(t, got, want)
	})
}

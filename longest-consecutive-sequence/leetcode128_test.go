package leetcode128

import (
	"testing"
)

func TestValidSudoku(t *testing.T) {
	assertInt := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		got := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
		want := 4
		assertInt(t, got, want)
	})

	t.Run("test 2", func(t *testing.T) {
		got := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
		want := 9
		assertInt(t, got, want)
	})
}

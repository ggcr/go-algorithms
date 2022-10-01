package leetcode36

import (
	"testing"
)

func TestValidSudoku(t *testing.T) {
	assertBoolean := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("double cumulative array", func(t *testing.T) {
		sudoku := [][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}
		got := isValidSudoku(sudoku)
		want := true
		assertBoolean(t, got, want)
	})
}

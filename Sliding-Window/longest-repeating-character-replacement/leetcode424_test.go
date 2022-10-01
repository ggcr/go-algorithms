package leetcode424

import (
	"reflect"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	assertSlice := func(t testing.TB, got, want int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test (1): \"ABAB\", 2", func(t *testing.T) {
		got := characterReplacement("ABAB", 2)
		want := 4
		assertSlice(t, got, want)
	})

	t.Run("test (2): \"AABABBA\", 1", func(t *testing.T) {
		got := characterReplacement("AABABBA", 1)
		want := 4
		assertSlice(t, got, want)
	})

	t.Run("test (3): \"AAAB\", 0", func(t *testing.T) {
		got := characterReplacement("AAAB", 0)
		want := 3
		assertSlice(t, got, want)
	})

	t.Run("test (4): \"AAAA\", 2", func(t *testing.T) {
		got := characterReplacement("AAAA", 2)
		want := 4
		assertSlice(t, got, want)
	})

	t.Run("test (5): \"ABBB\", 1", func(t *testing.T) {
		got := characterReplacement("ABBB", 1)
		want := 4
		assertSlice(t, got, want)
	})

	t.Run("test (6): \"ABCDE\", 1", func(t *testing.T) {
		got := characterReplacement("ABCDE", 1)
		want := 2
		assertSlice(t, got, want)
	})
}

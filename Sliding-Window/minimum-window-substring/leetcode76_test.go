package leetcode76

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assert := func(t testing.TB, got, want string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := minWindow("ADOBECODEBANC", "ABC")
		want := "BANC"
		assert(t, got, want)
	})

	t.Run("test (2)", func(t *testing.T) {
		got := minWindow("a", "a")
		want := "a"
		assert(t, got, want)
	})

	t.Run("test (3)", func(t *testing.T) {
		got := minWindow("a", "aa")
		want := ""
		assert(t, got, want)
	})

	t.Run("test (4)", func(t *testing.T) {
		got := minWindow("ab", "b")
		want := "b"
		assert(t, got, want)
	})

	t.Run("test (5)", func(t *testing.T) {
		got := minWindow("acbba", "aab")
		want := "acbba"
		assert(t, got, want)
	})

	t.Run("test (6)", func(t *testing.T) {
		got := minWindow("abc", "b")
		want := "b"
		assert(t, got, want)
	})
}

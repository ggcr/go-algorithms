package leetcode242

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	assertBoolean := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("HASH SET (1)", func(t *testing.T) {
		got := isAnagramSort("anagram", "nagaram")
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("HASH SET (2)", func(t *testing.T) {
		got := isAnagramSort("car", "rat")
		want := false
		assertBoolean(t, got, want)
	})

	t.Run("HASH SET (3)", func(t *testing.T) {
		got := isAnagramSort("ab", "ba")
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("HASH SET (1) SubAnagram", func(t *testing.T) {
		got := isSubAnagram("anagram", "nagaram")
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("HASH SET (2) SubAnagram", func(t *testing.T) {
		got := isSubAnagram("car", "rat")
		want := false
		assertBoolean(t, got, want)
	})

	t.Run("HASH SET (3) SubAnagram", func(t *testing.T) {
		got := isSubAnagram("ab", "ba")
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("SORT (1)", func(t *testing.T) {
		got := isAnagram("anagram", "nagaram")
		want := true
		assertBoolean(t, got, want)
	})

	t.Run("SORT (2)", func(t *testing.T) {
		got := isAnagram("car", "rat")
		want := false
		assertBoolean(t, got, want)
	})

	t.Run("SORT (3)", func(t *testing.T) {
		got := isAnagram("ab", "ba")
		want := true
		assertBoolean(t, got, want)
	})
}

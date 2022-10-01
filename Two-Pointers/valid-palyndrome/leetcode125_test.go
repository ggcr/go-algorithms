package leetcode125

import (
	"testing"
)

func TestStripAlphanumeric(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	got := strip(s)
	want := "amanaplanacanalpanama"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestValidPalindrome(t *testing.T) {
	t.Run("test leetcode", func(t *testing.T) {
		s := "A man, a plan, a canal: Panama"
		got := isPalindrome(s)
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("test even", func(t *testing.T) {
		s := "aaee"
		got := isPalindrome(s)
		want := false
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("test odd", func(t *testing.T) {
		s := "aae"
		got := isPalindrome(s)
		want := false
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("test weird", func(t *testing.T) {
		s := "0P"
		got := isPalindrome(s)
		want := false
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

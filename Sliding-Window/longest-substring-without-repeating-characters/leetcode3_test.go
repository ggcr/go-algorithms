package leetcode3

import (
	"reflect"
	"testing"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	assertSlice := func(t testing.TB, got, want int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := lengthOfLongestSubstring("abcabcbb")
		want := 3
		assertSlice(t, got, want)
	})

	t.Run("test (2)", func(t *testing.T) {
		got := lengthOfLongestSubstring("aa")
		want := 1
		assertSlice(t, got, want)
	})

	t.Run("test (3)", func(t *testing.T) {
		got := lengthOfLongestSubstring("au")
		want := 2
		assertSlice(t, got, want)
	})
}

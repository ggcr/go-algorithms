package leetcode567

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assert := func(t testing.TB, got, want bool) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := checkInclusion("ab", "eidbaooo")
		want := true
		assert(t, got, want)
	})

	t.Run("test (2)", func(t *testing.T) {
		got := checkInclusion("ab", "eidboaoo")
		want := false
		assert(t, got, want)
	})

	t.Run("test (3)", func(t *testing.T) {
		got := checkInclusion("adc", "dcda")
		want := true
		assert(t, got, want)
	})

	t.Run("test (4)", func(t *testing.T) {
		got := checkInclusion("abc", "bbbca")
		want := true
		assert(t, got, want)
	})

	t.Run("test (5)", func(t *testing.T) {
		got := checkInclusion("hello", "ooolleoooleh")
		want := false
		assert(t, got, want)
	})
}

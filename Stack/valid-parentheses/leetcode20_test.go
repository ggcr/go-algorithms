package leetcode20

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assert := func(t testing.TB, got, want bool) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("Test", func(t *testing.T) {
		got := isValid("()[]{}")
		want := true
		assert(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := isValid("()")
		want := true
		assert(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := isValid("(]")
		want := false
		assert(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := isValid("")
		want := false
		assert(t, got, want)
	})
}

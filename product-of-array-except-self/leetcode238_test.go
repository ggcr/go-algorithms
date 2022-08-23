package leetcode238

import (
	"reflect"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	assertSlice := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("double cumulative array", func(t *testing.T) {
		got := productExceptSelf([]int{-1, 1, 0, -3, 3})
		want := []int{0, 0, 9, 0, 0}
		assertSlice(t, got, want)
	})
}

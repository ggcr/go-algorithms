package leetcode121

import (
	"reflect"
	"testing"
)

func TestContainer(t *testing.T) {
	assertSlice := func(t testing.TB, got, want int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
		want := 49
		assertSlice(t, got, want)
	})
}

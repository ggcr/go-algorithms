package leetcode239

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assert := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
		want := []int{3, 3, 5, 5, 6, 7}
		assert(t, got, want)
	})
}

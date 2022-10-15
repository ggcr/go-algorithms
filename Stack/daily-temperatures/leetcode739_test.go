package leetcode739

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assert := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test", func(t *testing.T) {
		got := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
		want := []int{1, 1, 4, 2, 1, 1, 0, 0}
		assert(t, got, want)
	})
}

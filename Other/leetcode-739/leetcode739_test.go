package leetcode739

import (
	"reflect"
	"testing"
)

func Test_DailyTemperatures(t *testing.T) {
	assertEqualSlice := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		got := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
		want := []int{1, 1, 4, 2, 1, 1, 0, 0}
		assertEqualSlice(t, got, want)
	})

	t.Run("test 2", func(t *testing.T) {
		got := dailyTemperatures([]int{30, 40, 50, 60})
		want := []int{1, 1, 1, 0}
		assertEqualSlice(t, got, want)
	})

	t.Run("test 3", func(t *testing.T) {
		got := dailyTemperatures([]int{30, 60, 90})
		want := []int{1, 1, 0}
		assertEqualSlice(t, got, want)
	})

	// t.Run("test 4", func(t *testing.T) {
	// 	got := dailyTemperatures([]int{})
	// 	want := []int{}
	// 	assertEqualSlice(t, got, want)
	// })

}

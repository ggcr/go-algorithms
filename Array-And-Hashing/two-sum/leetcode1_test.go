package leetcode1

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assertSlice := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run(" (1)", func(t *testing.T) {
		got := twoSum([]int{2, 7, 11, 15}, 9)
		want := []int{0, 1}
		assertSlice(t, got, want)
	})

	t.Run(" (1)", func(t *testing.T) {
		got := twoSumMap([]int{2, 7, 11, 15}, 9)
		want := []int{0, 1}
		assertSlice(t, got, want)
	})
}

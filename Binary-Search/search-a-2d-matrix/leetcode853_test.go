package leetcode74

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

	t.Run("Test1", func(t *testing.T) {
		got := searchMatrix(
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 60},
			},
			3)
		want := true
		assert(t, got, want)
	})

	t.Run("Test2", func(t *testing.T) {
		got := searchMatrix(
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 60},
			},
			13)
		want := false
		assert(t, got, want)
	})

	t.Run("Test3", func(t *testing.T) {
		got := searchMatrix(
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 60},
			},
			23)
		want := true
		assert(t, got, want)
	})

	t.Run("Test4", func(t *testing.T) {
		got := searchMatrix(
			[][]int{
				[]int{1, 3, 5},
			},
			3)
		want := true
		assert(t, got, want)
	})

	t.Run("Test5", func(t *testing.T) {
		got := searchMatrix(
			[][]int{
				[]int{1, 3, 5, 7},
			},
			3)
		want := true
		assert(t, got, want)
	})
}

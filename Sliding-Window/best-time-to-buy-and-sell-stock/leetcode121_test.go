package leetcode121

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
		got := maxProfit([]int{7, 1, 5, 3, 6, 4})
		want := 5
		assertSlice(t, got, want)
	})

	t.Run("test (2)", func(t *testing.T) {
		got := maxProfit([]int{7, 6, 4, 3, 1})
		want := 0
		assertSlice(t, got, want)
	})
}

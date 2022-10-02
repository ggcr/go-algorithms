package main

import (
	"reflect"
	"testing"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	assertSlice := func(t testing.TB, got, want [][]int32) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := groupSort([]int32{3, 3, 1, 2, 1})
		want := [][]int32{
			{1, 2},
			{3, 2},
			{2, 1},
		}
		assertSlice(t, got, want)
	})

	t.Run("test (2)", func(t *testing.T) {
		got := groupSort([]int32{7, 12, 3})
		want := [][]int32{
			{12, 1},
			{7, 1},
			{3, 1},
		}
		assertSlice(t, got, want)
	})
}

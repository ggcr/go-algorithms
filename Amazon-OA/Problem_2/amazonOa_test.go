package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assertSlice := func(t testing.TB, got, want int32) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test (1)", func(t *testing.T) {
		got := getMaximumGreyness([]string{"1010", "0101", "1010"})
		want := int32(1)
		assertSlice(t, got, want)
	})
}

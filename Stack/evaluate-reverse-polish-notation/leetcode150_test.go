package leetcode150

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assertInt := func(t testing.TB, got, want int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Test", func(t *testing.T) {
		got := evalRPN([]string{"3", "10", "5", "+", "*"})
		want := 45
		assertInt(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := evalRPN([]string{"4", "13", "5", "/", "+"})
		want := 6
		assertInt(t, got, want)
	})

	t.Run("Test", func(t *testing.T) {
		got := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
		want := 22
		assertInt(t, got, want)
	})
}

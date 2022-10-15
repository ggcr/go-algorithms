package leetcode22

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assert := func(t testing.TB, got, want []string) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Test", func(t *testing.T) {
		got := generateParenthesis(3)
		want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
		assert(t, got, want)
	})
}

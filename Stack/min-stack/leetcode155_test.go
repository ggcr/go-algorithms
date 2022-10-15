package leetcode155

import (
	"fmt"
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
		obj := Constructor()
		obj.Push(-2) // minElement = -2
		obj.Push(0)  // minElement = -2
		obj.Push(-3) // minElement = -3
		fmt.Println(obj)
		obj.Pop() // minElement = -2
		fmt.Println(obj)
		got := obj.GetMin()
		want := -2
		assertInt(t, got, want)
	})
}

package stack

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

	assertBool := func(t testing.TB, got, want bool) {
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("Test push", func(t *testing.T) {
		s := &stack{}
		s.push(1)
		s.push(2)
		got := s.values
		want := []int{2, 1}
		assert(t, got, want)
	})

	t.Run("Test pop", func(t *testing.T) {
		s := &stack{}
		s.push(1)
		s.push(2)
		s.pop()
		got := s.values
		want := []int{1}
		assert(t, got, want)
	})

	t.Run("Test empty", func(t *testing.T) {
		s := &stack{}
		got := s.empty()
		want := true
		assertBool(t, got, want)
	})

	t.Run("Test not empty", func(t *testing.T) {
		s := &stack{}
		s.push(1)
		got := s.empty()
		want := false
		assertBool(t, got, want)
	})
}

package leetcode271

import (
	"reflect"
	"testing"
)

func TestValidSudoku(t *testing.T) {
	assertString := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertListString := func(t testing.TB, got, want []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("validate encode", func(t *testing.T) {
		strs := []string{"Hello", "World"}
		var c Codec
		got := c.Encode(strs)
		want := "Hello" + SECRET_KEY + "World"
		assertString(t, got, want)
	})

	t.Run("validate encode manual", func(t *testing.T) {
		strs := []string{"Hello", "World"}
		var c Codec
		got := c.EncodeManual(strs)
		want := "Hello" + SECRET_KEY + "World"
		assertString(t, got, want)
	})

	t.Run("validate decode", func(t *testing.T) {
		strs := "Hello" + SECRET_KEY + "World"
		var c Codec
		got := c.Decode(strs)
		want := []string{"Hello", "World"}
		assertListString(t, got, want)
	})

	t.Run("validate decode manual", func(t *testing.T) {
		strs := "Hello" + SECRET_KEY + "World"
		var c Codec
		got := c.DecodeManual(strs)
		want := []string{"Hello", "World"}
		assertListString(t, got, want)
	})

	t.Run("test both", func(t *testing.T) {
		strs := []string{"Hello", "World"}
		var c Codec
		got := c.Decode(c.Encode(strs))
		want := []string{"Hello", "World"}
		assertListString(t, got, want)
	})

	t.Run("test both manual", func(t *testing.T) {
		strs := []string{"Hello", "World"}
		var c Codec
		got := c.DecodeManual(c.EncodeManual(strs))
		want := []string{"Hello", "World"}
		assertListString(t, got, want)
	})
}

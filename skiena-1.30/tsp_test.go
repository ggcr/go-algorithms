// tests for lsp

package tsp

import (
	"math"
	"reflect"
	"testing"
)

// When comapring two float numbers
const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < float64EqualityThreshold
}

func Test_Utils(t *testing.T) {
	t.Run("euclidean distance", func(t *testing.T) {
		got := dist(Point{0, 0}, Point{5, 2})
		want := 5.39
		if almostEqual(got, want) {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func Test_NearestNeigh(t *testing.T) {
	assertIntSlice := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("circle", func(t *testing.T) {
		got := NearestNeigh(Circle, 0)
		want := []int{0, 1, 2, 3, 4, 5, 6, 7}
		assertIntSlice(t, got, want)
	})

	t.Run("horizontal line", func(t *testing.T) {
		got := NearestNeigh(Line_Horizontal, 3)
		want := []int{0, 1, 2, 3, 4, 5, 6}
		assertIntSlice(t, got, want)
	})

	t.Run("vertical line", func(t *testing.T) {
		got := NearestNeigh(Line_Veritcal, 3)
		want := []int{0, 1, 2, 3, 4, 5, 6}
		assertIntSlice(t, got, want)
	})
}

func Test_ClosestPair(t *testing.T) {
	assertIntSlice := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("circle", func(t *testing.T) {
		got := ClosestPair(Circle, 0)
		want := []int{7, 6, 5, 4, 3, 2, 1, 0}
		assertIntSlice(t, got, want)
	})

	t.Run("horizontal line", func(t *testing.T) {
		got := ClosestPair(Line_Horizontal, 3)
		want := []int{0, 1, 2, 3, 4, 5, 6}
		assertIntSlice(t, got, want)
	})

	t.Run("vertical line", func(t *testing.T) {
		got := ClosestPair(Line_Veritcal, 3)
		want := []int{0, 1, 2, 3, 4, 5, 6}
		assertIntSlice(t, got, want)
	})
}

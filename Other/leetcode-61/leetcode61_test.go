package leetcode61

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	values := []int{1, 2, 3, 4}
	head := CreateLinkedList(values)
	expected := ReadLinkedList(head)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("got %v want %v", values, expected)
	}
}

func TestLastToFirst(t *testing.T) {
	values := []int{1, 2, 3, 4}
	head := CreateLinkedList(values)
	head = lastToFirst(head)
	got := ReadLinkedList(head)

	want := []int{4, 1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRotateList(t *testing.T) {
	assertEqualValues := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test 1: rotate 2 times one by one", func(t *testing.T) {
		values := []int{1, 2, 3, 4}
		head := CreateLinkedList(values)
		head = rotateRightOneByOne(head, 2)
		got := ReadLinkedList(head)
		want := []int{3, 4, 1, 2}
		assertEqualValues(t, got, want)
	})

	t.Run("test 1: rotate 2 times", func(t *testing.T) {
		values := []int{1, 2, 3, 4}
		head := CreateLinkedList(values)
		head = rotateRight(head, 2)
		got := ReadLinkedList(head)
		want := []int{3, 4, 1, 2}
		assertEqualValues(t, got, want)
	})

	t.Run("test 2: rotate 12 times", func(t *testing.T) {
		values := []int{1, 2, 3, 4}
		head := CreateLinkedList(values)
		head = rotateRight(head, 12)
		got := ReadLinkedList(head)
		want := []int{1, 2, 3, 4}
		assertEqualValues(t, got, want)
	})
}

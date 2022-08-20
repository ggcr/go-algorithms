package linkedlist

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


// Leetcode problem 61- Rotate List
// https://leetcode.com/problems/rotate-list/

// Given the head of a linked list, rotate the list to the right by k places.

// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]

package leetcode61

// PROBLEM SPECIFICATION

// LinkedList implementation

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(numbers []int) *ListNode {
	prevNode := &ListNode{Val: numbers[0], Next: nil}
	head := prevNode
	for _, val := range numbers[1:] {
		currentNode := ListNode{val, nil}
		prevNode.Next = &currentNode
		prevNode = &currentNode
	}
	return head
}

// utils

func ReadLinkedList(head *ListNode) []int {
	currPtr := head
	res := []int{}
	for {
		if currPtr == nil {
			return res
		}
		res = append(res, currPtr.Val)
		currPtr = currPtr.Next
	}
}

// rotate linkedlist utils

func lastToFirst(head *ListNode) *ListNode {
	currPtr := head
	for {
		if currPtr.Next.Next == nil {
			currPtr.Next.Next = head
			head = currPtr.Next
			currPtr.Next = nil
			return head
		}
		currPtr = currPtr.Next
	}
}

func getLengthLinkedList(head *ListNode) int {
	currPtr := head
	res := 0
	for {
		if currPtr == nil {
			return res
		}
		res += 1
		currPtr = currPtr.Next
	}
}

// Slow version
// O(n^2)
func rotateRightOneByOne(head *ListNode, k int) *ListNode {
	lengthList := getLengthLinkedList(head)
	if lengthList == 0 || (lengthList == 0 && k == 0) {
		return head
	}
	k = k % lengthList
	for i := 0; i < k; i++ {
		aux := lastToFirst(head)
		head = aux
	}
	return head
}

// Improved version faster than 100% of Go submissions!
// O(n)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	currPtr := head
	count := 1
	prevArr := []*ListNode{}
	for {
		prevArr = append(prevArr, currPtr)

		if currPtr.Next == nil {
			if count == 1 || (count == 1 && k == 0) || k%count == 0 {
				return head
			}
			k = k % count
			nth := count - k - 1
			currPtr.Next = head
			head = prevArr[nth].Next
			prevArr[nth].Next = nil
			return head
		}

		currPtr = currPtr.Next
		count += 1
	}
}

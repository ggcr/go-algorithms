// LinkedList implementation

package linkedlist

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

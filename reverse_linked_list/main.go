// 206. Reverse Linked List
// Easy

// Given the head of a singly linked list, reverse the list, and return the reversed list.

// Example 1:

// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
// Example 2:

// Input: head = [1,2]
// Output: [2,1]
// Example 3:

// Input: head = []
// Output: []

// Constraints:

// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000

package main

import "fmt"

func main() {
	// e := &ListNode{5, nil}
	// d := &ListNode{4, e}
	c := &ListNode{3, nil}
	// c := &ListNode{3, d}
	b := &ListNode{2, c}
	a := &ListNode{1, b}
	reverseList(a)
	fmt.Println(a.toArray())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	if prev.Next == nil {
		return prev
	}
	curr := head.Next
	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}
	return prev
}

func setNilIf(v *interface{}) {
	*v = nil
}

func (l *ListNode) toArray() []int {
	temp := []int{}
	for l != nil {
		temp = append(temp, l.Val)
		l = l.Next
	}
	return temp
}

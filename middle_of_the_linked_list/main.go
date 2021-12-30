// 876. Middle of the Linked List
// Easy

// Given the head of a singly linked list, return the middle node of the linked list.

// If there are two middle nodes, return the second middle node.

// Example 1:

// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.
// Example 2:

// Input: head = [1,2,3,4,5,6]
// Output: [4,5,6]
// Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

// Constraints:

// The number of nodes in the list is in the range [1, 100].
// 1 <= Node.val <= 100

package main

import "fmt"

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	fmt.Println(middleNode(head).Val)
}

/**
* Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
 }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNodeSlowFast(head *ListNode) *ListNode {
	secondPtr := head
	for secondPtr != nil && secondPtr.Next != nil {
		head = head.Next
		secondPtr = secondPtr.Next.Next
	}
	return head
}

func middleNode(head *ListNode) *ListNode {
	first := head
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	head = first
	for i := 0; i < (length)/2; i++ {
		head = head.Next
	}
	return head
}

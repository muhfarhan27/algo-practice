// 116. Populating Next Right Pointers in Each Node
// Medium

// You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

// struct Node {
//   int val;
//   Node *left;
//   Node *right;
//   Node *next;
// }
// Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

// Initially, all next pointers are set to NULL.

// Example 1:

// Input: root = [1,2,3,4,5,6,7]
// Output: [1,#,2,3,#,4,5,6,7,#]
// Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
// Example 2:

// Input: root = []
// Output: []

package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	a := Node{Val: 1, Left: nil, Right: nil, Next: nil}
	b := Node{Val: 2, Left: nil, Right: nil, Next: nil}
	c := Node{Val: 3, Left: nil, Right: nil, Next: nil}
	d := Node{Val: 4, Left: nil, Right: nil, Next: nil}
	e := Node{Val: 5, Left: nil, Right: nil, Next: nil}
	f := Node{Val: 6, Left: nil, Right: nil, Next: nil}
	g := Node{Val: 7, Left: nil, Right: nil, Next: nil}

	a.Left = &b
	b.Left = &d
	b.Right = &e
	a.Right = &c
	c.Left = &f
	c.Right = &g
	aNew := connect(&a)
	print(aNew)
}

func connect(root *Node) *Node {
	leftNode := root
	for leftNode != nil && leftNode.Left != nil {
		helper(leftNode)
		leftNode = leftNode.Left
	}
	return root
}

func helper(head *Node) {
	iter := head
	for iter != nil {
		iter.Left.Next = iter.Right
		if iter.Next != nil {
			iter.Right.Next = iter.Next.Left
		}
		iter = iter.Next
	}
}

func print(root *Node) {
	fmt.Println("val:", root.Val)
	if root.Next != nil {
		fmt.Println("next:", root.Next.Val)
	} else {
		fmt.Println("next: nil")
	}
	if root.Left != nil {
		print(root.Left)
	}
	if root.Right != nil {
		print(root.Right)
	}
}

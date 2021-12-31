// 617. Merge Two Binary Trees

// You are given two binary trees root1 and root2.

// Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

// Return the merged tree.

// Note: The merging process must start from the root nodes of both trees.

// Example 1:

// Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// Output: [3,4,5,5,4,null,7]
// Example 2:

// Input: root1 = [1], root2 = [1,2]
// Output: [2,2]

package main

import "fmt"

func main() {
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node5 := &TreeNode{5, nil, nil}

	node1.Left = node3
	node1.Right = node2
	node3.Left = node5

	node1_2 := &TreeNode{1, nil, nil}
	node2_2 := &TreeNode{2, nil, nil}
	node3_2 := &TreeNode{3, nil, nil}
	node4_2 := &TreeNode{4, nil, nil}
	node7_2 := &TreeNode{7, nil, nil}

	node2_2.Left = node1_2
	node2_2.Right = node3_2
	node1_2.Right = node4_2
	node3_2.Right = node7_2

	print(mergeTrees(node1, node2_2))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func print(root *TreeNode) {
	fmt.Println(root.Val)
	if root.Left != nil {
		print(root.Left)
	}
	if root.Right != nil {
		print(root.Right)
	}
}

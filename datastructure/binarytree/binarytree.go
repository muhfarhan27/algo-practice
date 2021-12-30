package binarytree

import (
	queueCustom "datastructure/queue"
	stackCustom "datastructure/stack"
)

type BinaryNode struct {
	Value interface{}
	Left  *BinaryNode
	Right *BinaryNode
}

type BinaryTree struct {
	Root *BinaryNode
}

func (t *BinaryTree) DFS() []interface{} {
	stackInit := stackCustom.CustomStack{}
	stack := stackInit.New()
	temp := []interface{}{}
	if t.Root == nil {
		return temp
	}
	stack.Push(t.Root)
	for stack.Size() > 0 {
		if current, success := stack.Pop(); success {
			if node, ok := current.(*BinaryNode); ok {
				temp = append(temp, node.Value)
				//right first because stack LIFO
				if node.Right != nil {
					stack.Push(node.Right)
				}
				if node.Left != nil {
					stack.Push(node.Left)
				}
			}
		}
	}
	return temp
}

func (t *BinaryTree) BFS() []interface{} {
	qCustom := queueCustom.CustomQueue{}
	q := qCustom.New()
	temp := []interface{}{}
	if t.Root == nil {
		return temp
	}
	q.Enqueue(t.Root)
	for q.Size() > 0 {
		if current, success := q.Dequeue(); success {
			if node, ok := current.(*BinaryNode); ok {
				temp = append(temp, node.Value)
				if node.Left != nil {
					q.Enqueue(node.Left)
				}
				if node.Right != nil {
					q.Enqueue(node.Right)
				}
			}
		}
	}
	return temp
}

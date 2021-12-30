package main

import (
	binarytree "datastructure/binarytree"
	"fmt"
)

func main() {
	// 	   1
	// 	  / \
	// 	 2	 3
	//  / \	  \
	// 4   5   6
	one := binarytree.BinaryNode{Value: 1, Left: nil, Right: nil}
	two := binarytree.BinaryNode{Value: 2, Left: nil, Right: nil}
	three := binarytree.BinaryNode{Value: 3, Left: nil, Right: nil}
	four := binarytree.BinaryNode{Value: 4, Left: nil, Right: nil}
	five := binarytree.BinaryNode{Value: 5, Left: nil, Right: nil}
	six := binarytree.BinaryNode{Value: 6, Left: nil, Right: nil}

	one.Left = &two
	one.Right = &three
	two.Left = &four
	two.Right = &five
	three.Right = &six

	btree := binarytree.BinaryTree{Root: &one}
	fmt.Println(btree.DFS()) //[1 2 4 5 3 6]
	fmt.Println(btree.BFS()) //[1 2 3 4 5 6]
}

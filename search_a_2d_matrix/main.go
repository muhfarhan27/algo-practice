// 74. Search a 2D Matrix
// Medium

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the previous row.

// Example 1:

// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true
// Example 2:

// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	row := binarySerchGetRow(matrix, target, rows)
	if row < 0 {
		return false
	}
	return binarySearchPerRow(matrix, row, target, cols)
}

func binarySerchGetRow(matrix [][]int, target, rows int) int {
	top := 0
	bot := rows - 1
	row := 0
	for top <= bot {
		row = top + int((bot-top)/2)
		if target > matrix[row][len(matrix[0])-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bot = row - 1
		} else {
			break
		}
	}
	if top > bot {
		return -1
	}
	return row
}

func binarySearchPerRow(matrix [][]int, rowIndex, target, columns int) bool {
	l := 0
	r := columns - 1
	mid := 0
	for l <= r {
		mid = l + int((r-l)/2)
		if target == matrix[rowIndex][mid] {
			return true
		} else {
			if target > matrix[rowIndex][mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}

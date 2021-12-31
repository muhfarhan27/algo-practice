// 118. Pascal's Triangle
// Easy

// Given an integer numRows, return the first numRows of Pascal's triangle.

// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

// Example 1:

// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// Example 2:

// Input: numRows = 1
// Output: [[1]]

// Constraints:

// 1 <= numRows <= 30

package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	one := []int{1}
	two := []int{1, 1}
	res := [][]int{}
	res = append(res, one)
	if numRows == 1 {
		return res
	}
	res = append(res, two)
	if numRows == 2 {
		return res
	}
	prev := two
	for i := 3; i <= numRows; i++ {
		new := genFromPrev(i, prev)
		res = append(res, new)
		prev = new
	}

	return res
}

func genFromPrev(num int, prev []int) []int {
	res := make([]int, num)
	res[0] = prev[0]
	res[num-1] = prev[num-2]
	for i := 1; i < len(res)-1; i++ {
		res[i] = prev[i-1] + prev[i]
	}
	return res
}
